package keeper

import (
	"context"
	"crypto/sha256"
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
	if game.Challenger == msg.Creator {
		return nil, errors.New("cannot accept own game invitation")
	}

	// calculate player x and player o
	hsha := sha256.Sum256([]byte(game.Challenger + msg.Creator))
	lastBit := binary.BigEndian.Uint32(hsha[:]) & 1

	var playerX, playerO string
	if lastBit == 0 {
		playerX = msg.Creator
		playerO = game.Challenger
	} else {
		playerX = game.Challenger
		playerO = msg.Creator
	}

	// set new game data
	game.Status = types.GameStatus_IN_PROGRESS
	game.Opponent = msg.Creator
	game.PlayerX = playerX
	game.PlayerO = playerO

	// update store
	gameBytes, err = k.cdc.Marshal(&game)
	if err != nil {
		return nil, err
	}
	gamesStore.Set(gameKey, gameBytes)

	return &types.MsgAcceptGameResponse{}, nil
}
