package keeper

import (
	"context"
	"encoding/binary"

	"tic-tac-toe/x/tictactoe/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// retrieve the current game count to be used as the game ID
	gameCountStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.GameCountKey))
	// retrieve the games store for storing game data
	gamesStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.GamesKey))

	// get the current game count to use as ID for tictactoe game
	gameCountBytes := gameCountStore.Get([]byte(types.GameCountKey))
	currentGameCount := binary.BigEndian.Uint32(gameCountBytes)

	// create game struct
	game := types.Game{Id: currentGameCount}

	// store the game data
	gameBytes, err := game.Marshal()
	if err != nil {
		return nil, err
	}
	gamesStore.Set(gameCountBytes, gameBytes)

	// increment game count
	k.incrementGameCount(ctx)

	return &types.MsgCreateGameResponse{}, nil
}
