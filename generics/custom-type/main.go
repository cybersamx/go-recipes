package main

import (
	"fmt"
)

var (
	// This ensures that Cat and Dog conform to the behaviors of Animal.
	_ Animal = (*Cat)(nil)
	_ Animal = (*Dog)(nil)
)

type Animal interface {
	Species() string
	Talk() string
}

type Cat struct{}

func (c Cat) Species() string {
	return "cat"
}

func (c Cat) Talk() string {
	return "meow"
}

type Dog struct{}

func (d Dog) Species() string {
	return "dog"
}

func (d Dog) Talk() string {
	return "woof"
}

func Print[T Animal](animal T) {
	fmt.Println(animal.Talk())
}

func main() {
	pets := []Animal{
		Cat{},
		Dog{},
	}

	for _, p := range pets {
		Print(p)
	}
}
