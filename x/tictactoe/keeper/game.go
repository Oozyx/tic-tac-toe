package keeper

import (
	"encoding/binary"
	"tic-tac-toe/x/tictactoe/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetGameCount(ctx sdk.Context) uint32 {
	// retrieve the current game count store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.GameCountKey))
	// retrieve the current game count from the store
	gameCountBytes := store.Get([]byte(types.GameCountKey))

	// no entry yet so return 0
	if gameCountBytes == nil {
		return 0
	}

	return binary.BigEndian.Uint32(gameCountBytes)
}

func (k Keeper) incrementGameCount(ctx sdk.Context) {
	// retrieve the current game count store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.GameCountKey))
	// retrieve the current game count from the store
	gameCountBytes := store.Get([]byte(types.GameCountKey))
	gameCount := binary.BigEndian.Uint32(gameCountBytes)

	// increment
	newGameCountBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(newGameCountBytes, gameCount+1)
	store.Set([]byte(types.GameCountKey), newGameCountBytes)
}
