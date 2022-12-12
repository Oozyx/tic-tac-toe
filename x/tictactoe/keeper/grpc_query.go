package keeper

import (
	"tic-tac-toe/x/tictactoe/types"
)

var _ types.QueryServer = Keeper{}
