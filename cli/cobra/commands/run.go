package commands

import (
	"github.com/spf13/cobra"
)

var runCmd = cobra.Command{
	Use:   "run",                        // Name of the command
	Short: "Run a task on the cluster.", // Short descriptor in the usage table of cobra-cmd --help
	Long:  "Run a task on the cluster.", // Long help description in the output of cobra-cmd run --help
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	printArgs(cmd, args)

	return nil
}
