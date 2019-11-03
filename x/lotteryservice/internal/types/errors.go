package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeLotteryNotExist    sdk.CodeType = 101
	CodePermissionError    sdk.CodeType = 102
	CodeNeedMoreCandidates sdk.CodeType = 103
	CodeDoubleStart        sdk.CodeType = 104
)

// ErrNameDoesNotExist is the error for name not existing
func ErrLotteryDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeLotteryNotExist, "Lottery does not exist")
}

func ErrPermissionError(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodePermissionError, "You don't have permission to do this")
}

func ErrNeedMoreCandidates(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNeedMoreCandidates, "Too less candidates to start lottery")
}

func ErrDoubleStart(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeDoubleStart, "A lottery can't be start twice")
}
