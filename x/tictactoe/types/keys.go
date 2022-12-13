package types

const (
	// ModuleName defines the module name
	ModuleName = "tictactoe"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tictactoe"

	// GameCountKey defines the memory store for the amount of games created. Used as a unique game id.
	GameCountKey = "game/count"

	// GamesKey defines the memory store for the game data
	GamesKey = "game/games"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
