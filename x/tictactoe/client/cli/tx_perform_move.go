package cli

import (
	"strconv"

	"tic-tac-toe/x/tictactoe/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPerformMove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perform-move [game-id] [index]",
		Short: "Broadcast message PerformMove",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGameId, err := cast.ToUint32E(args[0])
			if err != nil {
				return err
			}
			argIndex, err := cast.ToUint32E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPerformMove(
				clientCtx.GetFromAddress().String(),
				argGameId,
				argIndex,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
