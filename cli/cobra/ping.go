package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func pingCommand(cfg *config, v *viper.Viper) *cobra.Command {
	// List command
	cmd := cobra.Command{
		Use:     "ping",
		Short:   "Ping the minio server",
		Example: "cobra ping",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Config:", cfg)

			return nil
		},
	}

	return &cmd
}
