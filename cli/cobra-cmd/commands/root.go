package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "cobra-cmd",
	Short: "cobra-cmd manages a cluster of machines.",
	Long:  "cobra-cmd manages a cluster of machines.",
	Run:   root,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func root(cmd *cobra.Command, _ []string) {
	// If no subcomand is entered, we default to the running the help subcommand.
	cobra.CheckErr(cmd.Help())
}
