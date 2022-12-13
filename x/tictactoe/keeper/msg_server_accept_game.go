package keeper

import (
	"context"
	"encoding/binary"
	"errors"

	"tic-tac-toe/x/tictactoe/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AcceptGame(goCtx context.Context, msg *types.MsgAcceptGame) (*types.MsgAcceptGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// initial validation that game id exists
	gameCount, _ := k.GetGameCount(ctx)
	if msg.GameId >= gameCount {
		return nil, errors.New("invalid game id")
	}

	// retrieve the games store for storing game data
	gamesStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.GamesKey))

	// retrieve the game invitation
	gameKey := make([]byte, 4)
	binary.BigEndian.PutUint32(gameKey, msg.GameId)
	gameBytes := gamesStore.Get(gameKey)
	var game types.Game
	err := k.cdc.Unmarshal(gameBytes, &game)
	if err != nil {
		return nil, err
	}

	// validate the game is still open
	if game.Status != types.GameStatus_OPEN {
		return nil, errors.New("game no longer open")
	}

	// validate sender is not accepting own game
	if game.PlayerOne == msg.Creator {
		return nil, errors.New("cannot accept own game invitation")
	}

	// set new game data
	game.Status = types.GameStatus_IN_PROGRESS
	game.PlayerTwo = msg.Creator

	// update store
	gameBytes, err = k.cdc.Marshal(&game)
	if err != nil {
		return nil, err
	}
	gamesStore.Set(gameKey, gameBytes)

	return &types.MsgAcceptGameResponse{}, nil
}
