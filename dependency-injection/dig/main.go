package main

import "github.com/cybersamx/go-recipes/dependency-injection/dig/pkg"

func main() {
	container, err := pkg.NewContainer()
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
