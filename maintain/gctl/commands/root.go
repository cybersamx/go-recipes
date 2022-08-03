package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "gor",
	Short: "Go-recipe control.",
	Long:  "Go-recipe control.",
	Run:   root,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func root(cmd *cobra.Command, _ []string) {
	// If no subcommand is entered, we default to the running the help subcommand.
	cobra.CheckErr(cmd.Help())
}
