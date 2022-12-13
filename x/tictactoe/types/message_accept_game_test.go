package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"tic-tac-toe/testutil/sample"
)

func TestMsgAcceptGame_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAcceptGame
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAcceptGame{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAcceptGame{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
