package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPerformMove = "perform_move"

var _ sdk.Msg = &MsgPerformMove{}

func NewMsgPerformMove(creator string, gameId uint32, index uint32) *MsgPerformMove {
	return &MsgPerformMove{
		Creator: creator,
		GameId:  gameId,
		Index:   index,
	}
}

func (msg *MsgPerformMove) Route() string {
	return RouterKey
}

func (msg *MsgPerformMove) Type() string {
	return TypeMsgPerformMove
}

func (msg *MsgPerformMove) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPerformMove) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPerformMove) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
