package lotteryservice

import (
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/keeper"
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper           = keeper.NewKeeper
	NewMsgCreateLottery = types.NewMsgCreateLottery
	NewAddCandidates    = types.NewMsgAddCandidates
	NewQuerier          = keeper.NewQuerier
	//NewMsgBuyName    = types.NewMsgBuyName
	//NewMsgSetName    = types.NewMsgSetName
	//NewMsgDeleteName = types.NewMsgDeleteName
	//NewWhois         = types.NewWhois
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper           = keeper.Keeper
	MsgCreateLottery = types.MsgCreateLottery
	Lottery          = types.Lottery
	Candidates       = types.Candidates
	//MsgSetName      = types.MsgSetName
	//MsgBuyName      = types.MsgBuyName
	//MsgDeleteName   = types.MsgDeleteName
	//QueryResResolve = types.QueryResResolve
	//QueryResNames   = types.QueryResNames
)
