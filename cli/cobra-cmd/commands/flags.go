package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func initConfig() {
	// Set up environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CYBER")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags - using both short and long posix flags
	rootCmd.PersistentFlags().StringP("format", "f", "yaml", "format of the output")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debugging")
	// Just long posix flag.
	rootCmd.PersistentFlags().String("key", "", "key to connect to cluster")
	cobra.CheckErr(viper.BindPFlag("format", rootCmd.PersistentFlags().Lookup("format")))
	cobra.CheckErr(viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug")))
	cobra.CheckErr(viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key")))
	viper.SetDefault("format", "yaml")
	viper.SetDefault("debug", false)
	viper.SetDefault("key", "")
}

func printArgs(cmd *cobra.Command, args []string) {
	fmt.Println("----- Dump of cli arguments -----")
	fmt.Println("command:", cmd.Use)

	fmt.Print("args:")
	for _, arg := range args {
		fmt.Printf(" %s", arg)
	}
	fmt.Println()

	fmt.Print("flag-names: ")
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		fmt.Printf("%s ", flag.Name)
	})
	fmt.Println()

	fmt.Print("flag-values: ")
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		fmt.Printf("%s ", flag.Value)
	})
	fmt.Println()
}
