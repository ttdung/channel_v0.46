package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/ttdung/channel_v0.46/x/channel/types"
)

var _ = strconv.Itoa(0)

func CmdFund() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fund [creatoraddr] [channelid] [cointo-partner] [hashcode] [numblock] [multisig-addr]",
		Short: "Broadcast message fund",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCreatoraddr := args[0]
			argChannelid := args[1]
			argHashcode := args[3]
			argMultisigAddr := args[5]

			cmd.Flags().Set(flags.FlagFrom, argMultisigAddr)
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var argCointoPartner []*sdk.Coin
			arg2 := strings.Split(args[2], ":")

			argCointoPartner = make([]*sdk.Coin, len(arg2))
			for i, coin := range arg2 {
				decCoin, err := sdk.ParseDecCoin(coin)
				if err != nil {
					return err
				}
				c, _ := sdk.NormalizeDecCoin(decCoin).TruncateDecimal()
				argCointoPartner[i] = &c
			}

			argNumblock, err := strconv.ParseUint(args[4], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgFund(
				clientCtx.GetFromAddress().String(),
				argCreatoraddr,
				argChannelid,
				argCointoPartner,
				argHashcode,
				argNumblock,
				argMultisigAddr,
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
