package lotteryservice

import (
	"fmt"
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgCreateLottery:
			return handleMsgCreateLottery(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreateLottery(ctx sdk.Context, keeper Keeper, msg types.MsgCreateLottery) sdk.Result {
	//TODO: only a group of people are allowed to create lottery?
	lottery := &types.Lottery{
		Rounds:       []int{},
		Title:        msg.Title,
		Description:  msg.Description,
		Owner:        msg.Owner,
		Hashed:       msg.Hashed,
		StopEnroll:   false,
		CurrentRound: -1,
	}
	lotteryID := keeper.CreateLottery(ctx, lottery)

	//TODO: 是否能够正常返回？
	return sdk.Result{Data: []byte(lotteryID)}
}