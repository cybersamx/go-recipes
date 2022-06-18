package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRunCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "run",                        // Name of the subcommand
		Short: "Run a task on the cluster.", // Short help descriptor in the usage table in the output of myapp --help
		Long:  "Enter author's name",        // Long help description in the output of myapp run --help
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(
				"run",
				args,
				viper.GetString("format"),
				viper.GetBool("debug"),
				viper.GetString("key"))
			return nil
		},
	}

	return &cmd
}
