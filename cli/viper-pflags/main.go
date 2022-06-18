package main

import (
	"github.com/spf13/viper"
)

func main() {
	app := NewApp()
	v := viper.GetViper()

	app.BindConfig(v)
	app.LoadConfig(v)

	app.Run()
}
