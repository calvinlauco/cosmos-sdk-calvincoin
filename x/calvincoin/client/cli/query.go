package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/calvinlauco/cosmos-sdk-calvincoin/x/calvincoin/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group calvincoin queries under a subcommand
	calvincoinQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	calvincoinQueryCmd.AddCommand(
		flags.GetCommands(
			// TODO: Add query Cmds
			GetCmdBalance(queryRoute, cdc),
			GetCmdTotalSupply(queryRoute, cdc),
		)...,
	)

	return calvincoinQueryCmd
}

// TODO: Add Query Commands
func GetCmdBalance(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "balances [address]",
		Short: "Query account balances",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/balances/%s", queryRoute, addr), nil)
			if err != nil {
				fmt.Printf("could not resolve account - %s %v\n", addr, err)
				return nil
			}

			var out types.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdTotalSupply(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "totalsupply",
		Short: "Query total supply",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/totalsupply", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not resolve total supply - %v\n", err)
				return nil
			}

			var out types.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
