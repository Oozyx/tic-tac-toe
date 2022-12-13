package keeper

import (
	"encoding/binary"
	"tic-tac-toe/x/tictactoe/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetGameCount(ctx sdk.Context) (uint32, []byte) {
	// retrieve the current game count store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.GameCountKey))
	// retrieve the current game count from the store
	gameCountBytes := store.Get([]byte(types.GameCountKey))

	// no entry yet so return 0
	if gameCountBytes == nil {
		gb := make([]byte, 4)
		binary.BigEndian.PutUint32(gb, 0)
		return 0, gb
	}

	return binary.BigEndian.Uint32(gameCountBytes), gameCountBytes
}

func (k Keeper) incrementGameCount(ctx sdk.Context) {
	// retrieve the current game count store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.GameCountKey))
	// retrieve the current game count from the store
	gameCountBytes := store.Get([]byte(types.GameCountKey))
	var gameCount uint32
	if gameCountBytes == nil {
		gameCount = 0
	} else {
		gameCount = binary.BigEndian.Uint32(gameCountBytes)
	}

	// increment
	newGameCountBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(newGameCountBytes, gameCount+1)
	store.Set([]byte(types.GameCountKey), newGameCountBytes)
}
