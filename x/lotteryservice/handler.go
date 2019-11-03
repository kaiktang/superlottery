package lotteryservice

import (
	"fmt"
	"github.com/TomKKlalala/superchainer/util"
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var logger = util.GetLogger("lotteryservice")

// NewHandler returns a handler for "lotteryservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgCreateLottery:
			logger.Info("receive msg MsgCreateLottery")
			return handleMsgCreateLottery(ctx, keeper, msg)
		case types.MsgAddCandidates:
			logger.Info("receive msg MsgAddCandidates")
			return handleAddCandidates(ctx, keeper, msg)
		case types.MsgStartLottery:
			return handleStartLottery(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized lotteryservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreateLottery(ctx sdk.Context, keeper Keeper, msg types.MsgCreateLottery) sdk.Result {
	//TODO: only a group of people are allowed to create lottery?
	lottery := &types.Lottery{
		Rounds:       msg.Rounds,
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

func handleAddCandidates(ctx sdk.Context, keeper Keeper, msg types.MsgAddCandidates) sdk.Result {
	logger.Info("进入handler: " + msg.ID)
	err := keeper.AddCandidates(ctx, msg.ID, msg.Candidates, msg.Sender)
	if err != nil {
		return sdk.Result{Code: err.Code(), Codespace: err.Codespace()}
	}
	return sdk.Result{}
}

//TODO: 待完成
func handleStartLottery(ctx sdk.Context, keeper Keeper, msg types.MsgStartLottery) sdk.Result {
	err := keeper.StartLottery(ctx, msg.ID, msg.Sender)
	if err != nil {
		return sdk.Result{Code: err.Code(), Codespace: err.Codespace()}
	}

	return sdk.Result{}
}
