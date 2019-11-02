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
	QueryLottery = "lottery"
)

var logger = util.GetLogger("keeper")

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		logger.Info("querier receive path: " + strings.Join(path, "*"))
		switch path[0] {
		case QueryLottery:
			return queryLottery(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown lotteryservice query endpoint")
		}
	}
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
