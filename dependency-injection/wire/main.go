package main

import "github.com/cybersamx/go-recipes/dependency-injection/wire/pkg"

func main() {
	hs, err := pkg.InitializeServer()
	if err != nil {
		panic(err)
	}

	err = hs.Start()
	if err != nil {
		panic(err)
	}
}
