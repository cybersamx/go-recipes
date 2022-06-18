package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Port        int
	Debug       bool
	PostgresURI string `mapstructure:"postgres-uri"`
}

type BindConfigOpts struct {
	FlagSet *pflag.FlagSet
	Args    []string
}

// BindConfig set up and binds configuration to app. The order of setting viper file, flags, and
// env vars at build time doesn't influence the viper load order at runtime. The order is:
// default values > viper file > environment variables > CLI arguments
// For details on load precedence see https://github.com/spf13/viper#why-viper
func (a *App) BindConfig(v *viper.Viper, set ...BindConfigOpts) {
	// Config flag
	v.SetConfigName("viper")
	v.SetConfigType("yaml")
	v.AddConfigPath(".") // Working directory

	// CLI flags
	if len(set) == 0 {
		set = []BindConfigOpts{
			{
				FlagSet: pflag.CommandLine,
				Args:    os.Args,
			},
		}
	}
	opts := set[0]

	opts.FlagSet.Bool("debug", false, "Enable debug")
	opts.FlagSet.Int32("port", 4000, "The TCP port to run the application")
	opts.FlagSet.String("postgres-uri", "postgresql://localhost", "The uri of PostgreSQL to where the application connects")

	if err := opts.FlagSet.Parse(opts.Args); err != nil {
		panic(fmt.Errorf("failed to parse arguments: %v", err))
	}
	if err := v.BindPFlags(opts.FlagSet); err != nil {
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
	_ = v.ReadInConfig()

	if err := v.UnmarshalExact(&a.config); err != nil {
		panic(fmt.Errorf("failed to umarshal parsed configurations to viper: %v", err))
	}
}
