package tictactoe_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "tic-tac-toe/testutil/keeper"
	"tic-tac-toe/testutil/nullify"
	"tic-tac-toe/x/tictactoe"
	"tic-tac-toe/x/tictactoe/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TictactoeKeeper(t)
	tictactoe.InitGenesis(ctx, *k, genesisState)
	got := tictactoe.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
