package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "investcli",
	Short: "investcli helps you get information about your investments",
	Long:  "investcli helps you get information about your investments. It currently integrates with Coinbase only.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error while executing investcli '%s'\n", err)
		os.Exit(1)
	}
}
