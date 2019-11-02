package keeper

import (
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)
import abci "github.com/tendermint/tendermint/abci/types"

const (
	QueryLottery = "lottery"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryLottery:
			return queryLottery(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

func queryLottery(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	lottery := keeper.GetLottery(ctx, path[0])

	if lottery == nil {
		return []byte{}, sdk.ErrUnknownRequest("Can't find lottery")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryLottery{Lottery: *lottery})
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
