package cmd

import (
	"investcli/coin"
	"investcli/coinbase"
	"investcli/cryptodotcom"

	"github.com/spf13/cobra"
)

var isDevelopment bool
var accountsCommand = &cobra.Command{
	Use:     "accounts",
	Aliases: []string{"accts"},
	Short:   "Fetch all balances from your Crypto Accounts. Only Coinbase and Crypto.com suppported.",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		coinbaseBalances := coinbase.Accounts(isDevelopment)
		cryptoDotComBalances := cryptodotcom.Accounts(isDevelopment)

		coin.CalculateProportionBetweenBalances(append(coinbaseBalances, cryptoDotComBalances...))
	},
}

func init() {
	accountsCommand.Flags().BoolVarP(&isDevelopment, "dev", "D", false, "If development third party APIs won't be called. Mock responses will be used.")
	rootCommand.AddCommand(accountsCommand)
}
