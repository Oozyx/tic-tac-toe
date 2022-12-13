package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAcceptGame = "accept_game"

var _ sdk.Msg = &MsgAcceptGame{}

func NewMsgAcceptGame(creator string, gameId uint32) *MsgAcceptGame {
	return &MsgAcceptGame{
		Creator: creator,
		GameId:  gameId,
	}
}

func (msg *MsgAcceptGame) Route() string {
	return RouterKey
}

func (msg *MsgAcceptGame) Type() string {
	return TypeMsgAcceptGame
}

func (msg *MsgAcceptGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAcceptGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
