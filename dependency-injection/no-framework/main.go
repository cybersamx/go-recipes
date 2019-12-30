package main

import (
	"github.com/cybersamx/go-recipes/dependency-injection/no-framework/pkg"
)

func main() {
	settings := pkg.NewSettings()

	db, err := pkg.OpenConnection(settings)
	if err != nil {
		panic(err)
	}

	datastore, err := pkg.NewDataStore(db)
	if err != nil {
		panic(err)
	}

	server := pkg.NewHTTPServer(settings, datastore)
	err = server.Start()
	if err != nil {
		panic(err)
	}
}
