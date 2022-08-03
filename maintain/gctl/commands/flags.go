package commands

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Type int

func initConfig() {
	// Set up environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CYBER")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags - using both short and long posix flags
	rootCmd.PersistentFlags().BoolP("dry-run", "d", false, "enable dry-run")
	cobra.CheckErr(viper.BindPFlag("dry-run", rootCmd.PersistentFlags().Lookup("dry-run")))
	viper.SetDefault("dry-run", false)
}
