package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewDescribeCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "describe",
		Short: "Describe a specific resource object.",
		Long:  "Describe a specific resource object.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(
				"describe",
				args,
				viper.GetString("format"),
				viper.GetBool("debug"),
				viper.GetString("key"))
			return nil
		},
	}

	return &cmd
}
