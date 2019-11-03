package keeper

import (
	"fmt"
	"github.com/TomKKlalala/superchainer/util"
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)
import abci "github.com/tendermint/tendermint/abci/types"

const (
	QueryLottery    = "lottery"
	QueryLotteries  = "lotteries"
	QueryCandidates = "candidates"
	QueryWinners    = "winners"
)

var logger = util.GetLogger("keeper")

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		logger.Info("querier receive path: " + strings.Join(path, "*"))
		switch path[0] {
		case QueryLottery:
			if len(path) > 2 && path[2] == QueryCandidates {
				return queryCandidates(ctx, path[1:], req, keeper)
			}
			if len(path) > 2 && path[2] == QueryWinners {
				return queryWinners(ctx, path[1:], req, keeper)
			}
			return queryLottery(ctx, path[1:], req, keeper)
		case QueryLotteries:
			return queryLotteries(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown lotteryservice query endpoint")
		}
	}
}

func queryWinners(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	lotteryID := path[0]
	winners := keeper.GetWinners(ctx, lotteryID)

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryWinners(winners))
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil
}

func queryLotteries(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var lotteries []types.Lottery

	lotteryItr := keeper.GetLotteryIterator(ctx)
	for ; lotteryItr.Valid(); lotteryItr.Next() {
		var lottery types.Lottery
		keeper.cdc.MustUnmarshalBinaryBare(lotteryItr.Value(), &lottery)
		lotteries = append(lotteries, lottery)
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryLotteries(lotteries))
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryLottery(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	lotteryID := path[0]
	fmt.Println(fmt.Sprintf("lottery id: %s", lotteryID))
	lottery := keeper.GetLottery(ctx, lotteryID)

	if lottery == nil {
		return []byte{}, sdk.ErrUnknownRequest("Can't find lottery")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryLottery{Lottery: *lottery})
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryCandidates(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	lotteryID := path[0]
	candidates := keeper.GetCandidates(ctx, lotteryID)

	res, err := codec.MarshalJSONIndent(keeper.cdc, candidates)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
