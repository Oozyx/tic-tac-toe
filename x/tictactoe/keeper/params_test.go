package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "tic-tac-toe/testutil/keeper"
	"tic-tac-toe/x/tictactoe/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TictactoeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
