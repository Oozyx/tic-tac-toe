package keeper

import (
	"context"

	"tic-tac-toe/x/tictactoe/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OpenGames(goCtx context.Context, req *types.QueryOpenGamesRequest) (*types.QueryOpenGamesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// retrieve the games store
	gamesStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.GamesKey))

	// paginate the games
	var games []*types.Game
	pageRes, err := query.Paginate(gamesStore, req.Pagination, func(key, value []byte) error {
		var game types.Game
		if err := k.cdc.Unmarshal(value, &game); err != nil {
			return err
		}

		if game.Status == types.GameStatus_OPEN {
			games = append(games, &game)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &types.QueryOpenGamesResponse{Games: games, Pagination: pageRes}, nil
}
