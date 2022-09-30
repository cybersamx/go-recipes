package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "cobra",
	Short: "cobra manages a cluster of machines.",
	Long:  "cobra manages a cluster of machines.",
	Run:   root,
}

func root(cmd *cobra.Command, _ []string) {
	// If no subcommand is entered, we default to the running the help subcommand.
	cobra.CheckErr(cmd.Help())
}
