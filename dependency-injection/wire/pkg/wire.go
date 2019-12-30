//+build wireinject

package pkg

import (
	"github.com/google/wire"
)

func InitializeServer() (*HTTPServer, error) {
	wire.Build(
		NewHTTPServer,
		NewDataStore,
		OpenConnection,
		NewSettings,
	)

	return nil, nil // Just return something to satisfy the compiler
}
