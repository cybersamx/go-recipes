package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type FuncSlice[T Number] []T

type ComposeSlice[T Number] struct {
	val []T
}

func (fs *FuncSlice[T]) filter(fn func(T) bool) *FuncSlice[T] {
	var r FuncSlice[T]

	for _, item := range *fs {
		if fn(item) {
			r = append(r, item)
		}
	}

	*fs = r

	return fs
}

func (cs *ComposeSlice[T]) filter(fn func(T) bool) *ComposeSlice[T] {
	var r []T

	for _, item := range cs.val {
		if fn(item) {
			r = append(r, item)
		}
	}

	cs.val = r

	return cs
}

// --- Main ---

func main() {
	// Using FuncSlice

	// Emulates: filtered = original.filter(v -> v > 0)
	fs := FuncSlice[int]{2, 5, -3, -9, 0, 10, -5, 12, 13, 21, 8}

	fs.filter(func(i int) bool {
		// Filter positive numbers
		return i > 0
	}).filter(func(i int) bool {
		// Filter even
		return (i % 2) == 0
	})

	fmt.Println(fs)

	// Using ComposeSlice

	sample := []int{2, 5, -3, -9, 0, 10, -5, 12, 13, 21, 8}
	cs := ComposeSlice[int]{sample}

	cs.filter(func(i int) bool {
		// Filter positive numbers
		return i > 0
	}).filter(func(i int) bool {
		// Filter even
		return (i % 2) == 0
	})

	fmt.Println(cs.val)
}
