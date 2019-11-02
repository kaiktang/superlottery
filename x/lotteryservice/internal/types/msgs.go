package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const RouterKey = ModuleName

type MsgCreateLottery struct {
	Rounds      []int          `json:"rounds"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Owner       sdk.AccAddress `json:"owner"`
	Hashed      bool           `json:"hashed"`
}

func NewMsgCreateLottery(title string, description string, owner sdk.AccAddress, rounds []int, hashed bool) *MsgCreateLottery {
	return &MsgCreateLottery{Rounds: rounds, Title: title, Description: description, Owner: owner, Hashed: hashed}
}

// Route should return the name of the module
func (msg MsgCreateLottery) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateLottery) Type() string { return "create_lottery" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateLottery) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Title) == 0 || len(msg.Description) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Description cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateLottery) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateLottery) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
