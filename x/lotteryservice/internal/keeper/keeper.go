package keeper

import (
	"github.com/TomKKlalala/superchainer/util"
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"sync"
)

var lock sync.Mutex

type Keeper struct {
	key sdk.StoreKey
	cdc *codec.Codec
}

func NewKeeper(key sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		key: key,
		cdc: cdc,
	}
}

func (k Keeper) GetNextLotteryID(ctx sdk.Context) string {
	lock.Lock()
	defer lock.Unlock()

	store := ctx.KVStore(k.key)
	bz := store.Get([]byte(types.LotteryIDKey))
	var id string
	if len(bz) == 0 {
		id = "0"
	} else {
		var num int
		num, err := strconv.Atoi(string(bz))
		if err != nil {
			logger.Error(err)
		}
		num++
		id = strconv.Itoa(num)
	}
	store.Set([]byte(types.LotteryIDKey), []byte(id))

	logger.Info("GetNextLotteryID: " + id)
	return id
}

func (k Keeper) CreateLottery(ctx sdk.Context, lottery *types.Lottery) string {
	store := ctx.KVStore(k.key)

	id := k.GetNextLotteryID(ctx)
	lottery.ID = id
	logger.Info("try to create: " + (types.LotteryPrefix + id) + "   " + (*lottery).String())
	store.Set([]byte(types.LotteryPrefix+id), k.cdc.MustMarshalBinaryBare(lottery))

	return id
}

func (k Keeper) IsLotteryPresent(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.key)

	return store.Has([]byte(types.LotteryPrefix + id))
}

func (k Keeper) GetLottery(ctx sdk.Context, id string) *types.Lottery {
	store := ctx.KVStore(k.key)

	logger.Info("try to get lottey: " + id)
	if !k.IsLotteryPresent(ctx, id) {
		return nil
	}

	bz := store.Get([]byte(types.LotteryPrefix + id))
	if len(bz) == 0 {
		return nil
	}

	var lottery types.Lottery
	k.cdc.MustUnmarshalBinaryBare(bz, &lottery)
	return &lottery
}

func (k Keeper) SetLottery(ctx sdk.Context, lottery *types.Lottery) {
	store := ctx.KVStore(k.key)
	store.Set([]byte(types.LotteryPrefix+lottery.ID), k.cdc.MustMarshalBinaryBare(lottery))
}

func (k Keeper) GetLotteryIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.key)
	return sdk.KVStorePrefixIterator(store, []byte(types.LotteryPrefix))
}

func (k Keeper) AddCandidates(ctx sdk.Context, lotteryID string, candidates types.Candidates, sender sdk.AccAddress) sdk.Error {
	lottery := k.GetLottery(ctx, lotteryID)
	if !lottery.Owner.Equals(sender) {
		return types.ErrPermissionError(types.DefaultCodespace)
	}
	store := ctx.KVStore(k.key)
	if lottery.Hashed {
		for i := 0; i < len(candidates); i++ {
			cid := strconv.Itoa(i + lottery.CandidateNum)
			// store hash hex
			store.Set([]byte(types.LotteryCandidatesPrefix+lotteryID+"_"+cid), []byte(util.Byte2Hex(util.Sha256([]byte(candidates[i])))))
		}
	} else {
		for i := 0; i < len(candidates); i++ {
			cid := strconv.Itoa(i + lottery.CandidateNum)
			store.Set([]byte(types.LotteryCandidatesPrefix+lotteryID+"_"+cid), []byte(candidates[i]))
		}
	}
	lottery.CandidateNum += len(candidates)

	// update candidate number
	k.SetLottery(ctx, lottery)
	return nil
}

func (k Keeper) GetCandidates(ctx sdk.Context, lotteryID string) types.Candidates {
	store := ctx.KVStore(k.key)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.LotteryCandidatesPrefix+lotteryID))

	var candidates types.Candidates
	for ; iterator.Valid(); iterator.Next() {
		candidates = append(candidates, string(iterator.Value()))
	}

	return candidates
}
