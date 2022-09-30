package commands

import (
	"github.com/spf13/cobra"
)

var describeCmd = cobra.Command{
	Use:   "describe",                              // Name of the command
	Short: "Describe a specific resource object.",  // Short descriptor in the usage table of cobra-cmd --help
	Long:  "Describe a specific resource object..", // Long help description in the output of cobra-cmd describe --help
	RunE:  describe,
}

func describe(cmd *cobra.Command, args []string) error {
	printArgs(cmd, args)

	return nil
}
