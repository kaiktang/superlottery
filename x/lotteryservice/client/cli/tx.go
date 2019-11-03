package cli

import (
	"github.com/TomKKlalala/superchainer/util"
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"strconv"
)
import sdk "github.com/cosmos/cosmos-sdk/types"

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Lotteryservice transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	nameserviceTxCmd.AddCommand(client.PostCommands(
		GetCmdCreateLottery(cdc),
	)...)

	return nameserviceTxCmd
}

func GetCmdCreateLottery(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create [title] [description] [rounds] [hashed]",
		Short: "create lottery",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			roundsStr := args[2]
			rounds := util.StringToArray(roundsStr, ",")
			hashedStr := args[3]
			hashed, err := strconv.ParseBool(hashedStr)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateLottery(args[0], args[1], cliCtx.GetFromAddress(), rounds, hashed)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
