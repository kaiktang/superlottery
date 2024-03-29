package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateLottery{}, "lotteryservice/MsgCreateLottery", nil)
	cdc.RegisterConcrete(MsgAddCandidates{}, "lotteryservice/MsgAddCandidates", nil)
	cdc.RegisterConcrete(MsgStartLottery{}, "lotteryservice/MsgStartLottery", nil)
}
