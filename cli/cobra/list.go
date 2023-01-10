package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func listCommand(cfg *config, v *viper.Viper) *cobra.Command {
	// List command
	cmd := cobra.Command{
		Use:     "list",
		Short:   "List resources on the server",
		Example: "cobra list <buckets|objects>",
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var path string
			if len(args) > 0 {
				path = args[0]
			}

			fmt.Println("list args:", path)
			fmt.Println("Config:", cfg)

			return nil
		},
	}

	// Command list flags don't get extended to environment variables, so the
	// following is sufficient.
	flags := cmd.Flags()
	flags.BoolVarP(&cfg.recursive, "recursive", "r", false, "show objects recursively")
	flags.BoolVarP(&cfg.longFormat, "long", "l", true, "show the output in long format")
	checkErr(v.BindPFlags(flags))

	return &cmd
}
