package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeLotteryNotExist sdk.CodeType = 101
)

// ErrNameDoesNotExist is the error for name not existing
func ErrLotteryDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeLotteryNotExist, "Name does not exist")
}