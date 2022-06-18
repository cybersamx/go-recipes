package main

import (
	"strings"

	"github.com/cybersamx/go-recipes/config/cobra-cmd/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd *cobra.Command
)

func initConfig() {
	// Set up environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CYBER")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

func init() {
	cobra.OnInitialize(initConfig)

	// Root command
	rootCmd = commands.NewRootCommand()

	// Global flags
	rootCmd.PersistentFlags().String("format", "yaml", "format of the output")
	rootCmd.PersistentFlags().Bool("debug", false, "enable debugging")
	rootCmd.PersistentFlags().String("key", "", "key to connect to cluster")
	cobra.CheckErr(viper.BindPFlag("format", rootCmd.PersistentFlags().Lookup("format")))
	cobra.CheckErr(viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug")))
	cobra.CheckErr(viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key")))
	viper.SetDefault("format", "yaml")
	viper.SetDefault("debug", false)
	viper.SetDefault("key", "")

	// Subcommands
	rootCmd.AddCommand(commands.NewRunCommand())
	rootCmd.AddCommand(commands.NewDescribeCommand())
}

func main() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
