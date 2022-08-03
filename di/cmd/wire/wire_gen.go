// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/cybersamx/go-recipes/di/pkg"
)

// Injectors from wire.go:

func InitializeServer() (*pkg.HTTPServer, error) {
	settings := pkg.NewSettings()
	db, err := pkg.OpenConnection(settings)
	if err != nil {
		return nil, err
	}
	dataStore, err := pkg.NewDataStore(db)
	if err != nil {
		return nil, err
	}
	httpServer := pkg.NewHTTPServer(settings, dataStore)
	return httpServer, nil
}
