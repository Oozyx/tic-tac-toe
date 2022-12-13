package keeper

import (
	"context"

	"tic-tac-toe/x/tictactoe/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// retrieve the games store for storing game data
	gamesStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.GamesKey))

	// get the current game count to use as ID for tictactoe game
	currentGameCount, gameCountBytes := k.GetGameCount(ctx)

	// create game struct
	board := types.InitializeBoard()
	game := types.Game{Id: currentGameCount, Status: types.GameStatus_OPEN, Challenger: msg.Creator, Board: board}

	// store the game data
	gameBytes, err := game.Marshal()
	if err != nil {
		return nil, err
	}
	gamesStore.Set(gameCountBytes, gameBytes)

	// increment game count
	k.incrementGameCount(ctx)

	return &types.MsgCreateGameResponse{GameId: currentGameCount}, nil
}
