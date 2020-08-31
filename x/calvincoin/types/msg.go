package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgTransfer struct {
	From   sdk.AccAddress `json:"from"`
	To     sdk.AccAddress `json:"to"`
	Amount sdk.Coins      `json:"amount"`
}

func NewMsgTransfer(from sdk.AccAddress, to sdk.AccAddress, amount sdk.Coins) MsgTransfer {
	return MsgTransfer{
		From:   from,
		To:     to,
		Amount: amount,
	}
}

// Route should return the name of the module
func (msg MsgTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgTransfer) Type() string { return "transfer" }

// ValidateBasic runs stateless checks on the message
func (msg MsgTransfer) ValidateBasic() error {
	if msg.From.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.From.String())
	}
	if msg.To.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.To.String())
	}
	if !msg.Amount.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`
/*
// verify interface at compile time
var _ sdk.Msg = &Msg<Action>{}

// Msg<Action> - struct for unjailing jailed validator
type Msg<Action> struct {
	ValidatorAddr sdk.ValAddress `json:"address" yaml:"address"` // address of the validator operator
}

// NewMsg<Action> creates a new Msg<Action> instance
func NewMsg<Action>(validatorAddr sdk.ValAddress) Msg<Action> {
	return Msg<Action>{
		ValidatorAddr: validatorAddr,
	}
}

const <action>Const = "<action>"

// nolint
func (msg Msg<Action>) Route() string { return RouterKey }
func (msg Msg<Action>) Type() string  { return <action>Const }
func (msg Msg<Action>) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg Msg<Action>) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg Msg<Action>) ValidateBasic() error {
	if msg.ValidatorAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address")
	}
	return nil
}
*/
