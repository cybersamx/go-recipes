package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Debug       bool
	Port        int
	PostgresURL string `mapstructure:"postgres-url"`
	Directory   string
	Name        string
}

// BindConfig set up and binds configuration to app. The order of setting viper file, flags, and
// env vars at build time doesn't influence the viper load order at runtime. The order is:
// default values > viper file > environment variables > CLI arguments
// For details on load precedence see https://github.com/spf13/viper#why-viper
func (a *App) BindConfig(v *viper.Viper) {
	// Config flag
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".") // Working directory

	// CLI flags
	flagset := pflag.CommandLine
	flagset.Bool("debug", false, "debug")
	flagset.Int32("port", 4000, "port")
	flagset.String("postgres-url", "default", "postgres-url")
	flagset.String("directory", "default", "directory")
	flagset.String("name", "default", "name")

	if err := flagset.Parse(os.Args); err != nil {
		panic(fmt.Errorf("failed to parse arguments: %v", err))
	}
	if err := v.BindPFlags(flagset); err != nil {
		panic(fmt.Errorf("failed to bind pflags: %v", err))
	}

	// Environment variables
	// The setup allows the following mappings of env vars and flags (key)
	// CYBER_PORT         <--> port
	// CYBER_POSTGRES_URI <--> postgres-uri
	v.AutomaticEnv()
	v.SetEnvPrefix("CYBER")
	replacer := strings.NewReplacer("-", "_")
	v.SetEnvKeyReplacer(replacer)
}

// LoadConfig loads the configurations at runtime.
func (a *App) LoadConfig(v *viper.Viper) {
	// If there's no viper file to load, it's ok and move on.
	v.ReadInConfig()

	if err := v.UnmarshalExact(&a.config); err != nil {
		panic(fmt.Errorf("failed to umarshal parsed configurations to viper: %v", err))
	}
}
