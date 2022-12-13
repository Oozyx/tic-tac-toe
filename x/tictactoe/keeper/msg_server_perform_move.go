package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"tic-tac-toe/x/tictactoe/types"
)

func (k msgServer) PerformMove(goCtx context.Context, msg *types.MsgPerformMove) (*types.MsgPerformMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPerformMoveResponse{}, nil
}
