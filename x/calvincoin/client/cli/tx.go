package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/calvinlauco/cosmos-sdk-calvincoin/x/calvincoin/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	calvincoinTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	calvincoinTxCmd.AddCommand(flags.PostCommands(
		// TODO: Add tx based commands
		// GetCmd<Action>(cdc)
		GetCmdTransfer(cdc),
	)...)

	return calvincoinTxCmd
}

// Example:
//
// GetCmd<Action> is the CLI command for doing <Action>
// func GetCmd<Action>(cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "/* Describe your action cmd */",
// 		Short: "/* Provide a short description on the cmd */",
// 		Args:  cobra.ExactArgs(2), // Does your request require arguments
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)
// 			inBuf := bufio.NewReader(cmd.InOrStdin())
// 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

// 			msg := types.NewMsg<Action>(/* Action params */)
// 			err = msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}

// 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
// 		},
// 	}
// }
// GetCmdBuyName is the CLI command for sending a Transfer transaction
func GetCmdTransfer(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "transfer [address] [amount]",
		Short: "send CalvinCoin to another account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			amount, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			toAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgTransfer(cliCtx.GetFromAddress(), toAddr, amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
