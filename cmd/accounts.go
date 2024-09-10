package cmd

import (
	"investcli/coinbase"

	"github.com/spf13/cobra"
)

var isDevelopment bool
var accountsCommand = &cobra.Command{
	Use:     "accounts",
	Aliases: []string{"accts"},
	Short:   "Fetch all accounts associated with your Coinbase account",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		coinbase.Accounts(isDevelopment)
	},
}

func init() {
	accountsCommand.Flags().BoolVarP(&isDevelopment, "dev", "D", false, "If development third party APIs won't be called. Mock responses will be used.")
	rootCommand.AddCommand(accountsCommand)
}
