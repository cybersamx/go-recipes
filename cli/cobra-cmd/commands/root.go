package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = cobra.Command{
	Use:   "myctl",
	Short: "myctl manages my cluster of machines.",
	Long:  "myctl manages my cluster of machines.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("myctl")
	},
}

func initConfig() {
	// Set up environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CYBER")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

func init() {
	cobra.OnInitialize(initConfig)

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
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
