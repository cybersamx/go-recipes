package main

import (
	"fmt"
)

type Vector[T any] struct {
	elems []T
}

// There's no type parameter for a method. Its type parameter is defined at the struct.

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
