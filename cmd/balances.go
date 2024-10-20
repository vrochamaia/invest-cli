package cmd

import (
	"investcli/coin"
	"investcli/coinbase"
	"investcli/cryptodotcom"

	"github.com/spf13/cobra"
)

var isDevelopment bool
var balancesCommand = &cobra.Command{
	Use:     "balances",
	Aliases: []string{"balances"},
	Short:   "Fetch available balances from your Crypto Accounts and calculate weigth of which coin regarding total value of assets.",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		coinbaseBalances := coinbase.Balances(isDevelopment)
		cryptoDotComBalances := cryptodotcom.Balances(isDevelopment)

		coin.CalculateProportionAmongBalances(append(coinbaseBalances, cryptoDotComBalances...))
	},
}

func init() {
	balancesCommand.Flags().BoolVarP(&isDevelopment, "dev", "D", false, "If development third party APIs won't be called. Mock responses will be used.")
	rootCommand.AddCommand(balancesCommand)
}
