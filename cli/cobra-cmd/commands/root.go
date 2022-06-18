package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "myctl",
		Short: "myctl manages my cluster of machines.",
		Long:  "myctl manages my cluster of machines.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("app")
		},
	}

	return &cmd
}
