package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeLotteryNotExist sdk.CodeType = 101
	CodePermissionError sdk.CodeType = 102
)

// ErrNameDoesNotExist is the error for name not existing
func ErrLotteryDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeLotteryNotExist, "Lottery does not exist")
}

func ErrPermissionError(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodePermissionError, "You don't have permission to do this")
}
