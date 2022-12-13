package keeper

import (
	"context"
	"encoding/binary"
	"errors"

	"tic-tac-toe/x/tictactoe/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PerformMove(goCtx context.Context, msg *types.MsgPerformMove) (*types.MsgPerformMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// initial validation that game id exists
	gameCount, _ := k.GetGameCount(ctx)
	if msg.GameId >= gameCount {
		return nil, errors.New("invalid game id")
	}

	// retrieve the games store for storing game data
	gamesStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.GamesKey))

	// retrieve the game
	gameKey := make([]byte, 4)
	binary.BigEndian.PutUint32(gameKey, msg.GameId)
	gameBytes := gamesStore.Get(gameKey)
	var game types.Game
	err := k.cdc.Unmarshal(gameBytes, &game)
	if err != nil {
		return nil, err
	}

	// validate the game is in progress
	if game.Status != types.GameStatus_IN_PROGRESS {
		return nil, errors.New("game not in progress")
	}

	// validate the next turn
	if game.NextTurn != msg.Creator {
		return nil, errors.New("not creator's turn")
	}

	// validate index is empty
	if game.Board[msg.Index] != types.BoardEntry_EMPTY {
		return nil, errors.New("board index not empty")
	}

	// perform move
	var boardEntry types.BoardEntry
	if msg.Creator == game.PlayerX {
		boardEntry = types.BoardEntry_X
	} else {
		boardEntry = types.BoardEntry_O
	}
	game.Board[msg.Index] = boardEntry

	// toggle next turn
	if game.NextTurn == game.PlayerX {
		game.NextTurn = game.PlayerO
	} else {
		game.NextTurn = game.PlayerX
	}

	// check for completion and winner
	if game.IsGameComplete() || game.IsWinner(boardEntry) {
		game.Status = types.GameStatus_COMPLETE
	}
	if game.IsWinner(boardEntry) {
		game.Winner = msg.Creator
	}

	// update store
	gameBytes, err = k.cdc.Marshal(&game)
	if err != nil {
		return nil, err
	}
	gamesStore.Set(gameKey, gameBytes)

	return &types.MsgPerformMoveResponse{}, nil
}
