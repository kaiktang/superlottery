package keeper

import (
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

func (k Keeper) CreateLottery(ctx sdk.Context, lottery *types.Lottery) string {
	store := ctx.KVStore(k.key)

	id := types.LotteryID()
	store.Set([]byte(types.LotteryPrefix+id), k.cdc.MustMarshalBinaryBare(lottery))

	return id
}

func (k Keeper) IsLotteryPresent(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.key)

	return store.Has([]byte(types.LotteryPrefix + id))
}

func (k Keeper) GetLottery(ctx sdk.Context, id string) *types.Lottery {
	store := ctx.KVStore(k.key)

	if k.IsLotteryPresent(ctx, id) {
		return nil
	}

	bz := store.Get([]byte(types.LotteryPrefix + id))

	var lottery types.Lottery
	k.cdc.MustUnmarshalBinaryBare(bz, &lottery)
	return &lottery
}
