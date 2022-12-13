package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"tic-tac-toe/x/tictactoe/types"
)

func (k msgServer) AcceptGame(goCtx context.Context, msg *types.MsgAcceptGame) (*types.MsgAcceptGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAcceptGameResponse{}, nil
}
