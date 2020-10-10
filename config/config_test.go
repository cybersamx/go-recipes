package main

import (
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	appName = "cmd" // Something arbitrary
)

func getTestFlagSet() *pflag.FlagSet {
	return pflag.NewFlagSet(appName, pflag.ContinueOnError)
}

// CREDITS:
// Based on https://github.com/spf13/viper/blob/master/internal/testutil/env.go
// Licensed under the MIT license
// Copyright (c) 2017 Canonical Ltd.

func Setenv(t *testing.T, name, val string) {
	setenv(t, name, val, true)
}

func Unsetenv(t *testing.T, name string) {
	setenv(t, name, "", false)
}

func setenv(t *testing.T, name, val string, valOK bool) {
	oldVal, oldOK := os.LookupEnv(name)
	if valOK {
		require.NoError(t, os.Setenv(name, val))
	} else {
		require.NoError(t, os.Unsetenv(name))
	}

	// t.Cleanup is only available in Go 1.14+.
	t.Cleanup(func() {
		if oldOK {
			require.NoError(t, os.Setenv(name, oldVal))
		} else {
			require.NoError(t, os.Unsetenv(name))
		}
	})
}

func writeYAML(t *testing.T) {
	content := `
port: 5000
debug: true
postgres-uri: postgresql://postgres.us-east-2.rds.amazonaws.com/mydb
`

	filename := "config.yaml"
	file, err := os.Create(filename)
	require.NoError(t, err)

	defer func() {
		require.NoError(t, file.Close())
	}()

	_, err = file.Write([]byte(content))
	require.NoError(t, err)
}

func TestDefaultValues(t *testing.T) {
	viper.Reset()

	app := NewApp()
	v := viper.GetViper()

	opts := BindConfigOpts{
		FlagSet: getTestFlagSet(),
		Args:    []string{},
	}
	app.BindConfig(v, opts)
	app.LoadConfig(v)

	assert.Equal(t, 4000, app.config.Port)
	assert.Equal(t, false, app.config.Debug)
	assert.Equal(t, "postgresql://localhost", app.config.PostgresURI)
}

func TestLoadFromConfigFile(t *testing.T) {
	viper.Reset()

	app := NewApp()
	v := viper.GetViper()

	writeYAML(t)
	defer func() {
		require.NoError(t, os.Remove("config.yaml"))
	}()

	opts := BindConfigOpts{
		FlagSet: getTestFlagSet(),
		Args:    []string{},
	}
	app.BindConfig(v, opts)
	app.LoadConfig(v)

	assert.Equal(t, 5000, app.config.Port)
	assert.Equal(t, true, app.config.Debug)
	assert.Equal(t, "postgresql://postgres.us-east-2.rds.amazonaws.com/mydb", app.config.PostgresURI)
}

func TestLoadFromEnvVariables(t *testing.T) {
	viper.Reset()

	app := NewApp()
	v := viper.GetViper()

	Setenv(t, "CYBER_PORT", "6000")
	Setenv(t, "CYBER_DEBUG", "true")
	Setenv(t, "CYBER_POSTGRES_URI", "postgresql://10.20.30.40")
	defer func() {
		Unsetenv(t, "CYBER_PORT")
		Unsetenv(t, "CYBER_DEBUG")
		Unsetenv(t, "CYBER_POSTGRES_URI")
	}()

	opts := BindConfigOpts{
		FlagSet: getTestFlagSet(),
		Args:    []string{},
	}
	app.BindConfig(v, opts)
	app.LoadConfig(v)

	assert.Equal(t, 6000, app.config.Port)
	assert.Equal(t, true, app.config.Debug)
	assert.Equal(t, "postgresql://10.20.30.40", app.config.PostgresURI)
}

func TestFromLoadFlags(t *testing.T) {
	viper.Reset()

	app := NewApp()
	v := viper.GetViper()

	opts := BindConfigOpts{
		FlagSet: getTestFlagSet(),
		Args: []string{
			"--debug", "true",
			"--port", "7000",
			"--postgres-uri", "postgresql://pq.example.com:30000",
		},
	}
	app.BindConfig(v, opts)
	app.LoadConfig(v)

	assert.Equal(t, 7000, app.config.Port)
	assert.Equal(t, true, app.config.Debug)
	assert.Equal(t, "postgresql://pq.example.com:30000", app.config.PostgresURI)
}
