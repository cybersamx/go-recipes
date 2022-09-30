package main

import (
	"fmt"
)

type App struct {
	config Config
}

func (a *App) Run() {
	fmt.Println("Dump of config:", a.config)
}

func NewApp() *App {
	app := new(App)

	return app
}
