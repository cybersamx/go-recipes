//+build wireinject

package main

import (
	"github.com/cybersamx/go-recipes/di/pkg"
	"github.com/google/wire"
)

func InitializeServer() (*pkg.HTTPServer, error) {
	wire.Build(
		pkg.NewHTTPServer,
		pkg.NewDataStore,
		pkg.OpenConnection,
		pkg.NewSettings,
	)

	return nil, nil // Just return something to satisfy the compiler
}
