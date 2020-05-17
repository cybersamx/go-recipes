package main

import (
	"github.com/cybersamx/go-recipes/di/pkg"
)

func main() {
	container, err := NewContainer()
	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(hs *pkg.HTTPServer) {
		hs.Start()
	})
	if err != nil {
		panic(err)
	}
}
