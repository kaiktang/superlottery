package cli

import (
	"fmt"
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	lotteryserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the lotteryservice module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	lotteryserviceQueryCmd.AddCommand(client.GetCommands(
		GetCmdGetLottery(storeKey, cdc),
	)...)
	return lotteryserviceQueryCmd
}

func GetCmdGetLottery(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-lottery [id]",
		Short: "get lottery by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/lottery/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("could not find lottery: %s \n", id)
				return nil
			}

			var out types.QueryLottery
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
