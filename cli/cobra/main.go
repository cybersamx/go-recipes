package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "program exited due to %v", err)
		os.Exit(1)
	}
}

func rootCommand() *cobra.Command {
	cfg := config{}

	// Root command.
	cmd := cobra.Command{
		Use: "cobra",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// Use viper to manage environment variables and args.
	v := viper.New()
	v.SetEnvPrefix("CY")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Order of overrides
	// Default < environment variables < flags.

	// Assign defaults.
	cfg.addr = "localhost:9000"
	cfg.keyID = "admin"
	cfg.secretKey = "password"
	cfg.useSSL = false

	// Bind the flags to config object.
	flags := cmd.PersistentFlags()
	flags.StringVarP(&cfg.addr, "addr", "a", "localhost:9000", "the server address")
	flags.StringVar(&cfg.keyID, "key-id", "admin", "Access key id to access the server")
	flags.StringVar(&cfg.secretKey, "secret-key", "password", "Secret access key to access the server")
	flags.BoolVar(&cfg.useSSL, "use-ssl", false, "Use SSL to connect to the server")
	checkErr(v.BindPFlags(flags))

	// Get* is needed if we want to get the values from environment.
	cfg.addr = v.GetString("addr")
	cfg.keyID = v.GetString("key-id")
	cfg.secretKey = v.GetString("secret-key")
	cfg.useSSL = v.GetBool("use-ssl")

	// CLI commands.
	cmd.AddCommand(pingCommand(&cfg, v))
	cmd.AddCommand(listCommand(&cfg, v))

	return &cmd
}

func main() {
	err := rootCommand().Execute()
	checkErr(err)
}
