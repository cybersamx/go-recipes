package main

import (
	"log"
)

type App struct {
	config Config
}

func (a *App) Run() {
	log.Println("Running...")
	log.Println("viper", a.config)
}

func NewApp() *App {
	app := &App{}

	return app
}
