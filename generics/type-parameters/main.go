package main

import (
	"fmt"
)

type Vector[T any] struct {
	elems []T
}

// The method doesn't use a type parameter, instead the type parameter needs to be defined
// at the struct level.

func (v *Vector[T]) Push(e T) {
	v.elems = append(v.elems, e)
}

func main() {
	vi := Vector[int]{}
	vi.Push(1)
	vi.Push(2)
	fmt.Println(vi)

	vs := Vector[string]{}
	vs.Push("hello")
	vs.Push("world")
	fmt.Println(vs)
}
