package main

import "fmt"

type State struct {
	name         string
	abbreviation string
	population   int
}

func main() {
	fmt.Println("\n--- Slice ---")

	// Slice is nil.
	// If you want to instantiate and initialize a slice at the same time, do the following:
	// var states []state{}
	var states []State

	// For slice, initialization is optional. You can add an item to a nil slice.

	// --- Add an item ---
	ca := State{
		name:         "California",
		abbreviation: "CA",
		population:   39512223,
	}
	states = append(states, ca)
	fmt.Println(states)

	// --- Get an item ---

	state := states[0]
	fmt.Println(state)

	// Panics if index is out of bound.

	// --- Iterate the map ---

	ma := State{
		name:         "Massachusetts",
		abbreviation: "MA",
		population:   6892503,
	}
	states = append(states, ma)

	for i, item := range states {
		fmt.Printf("index: %d, item: %v\n", i, item)
	}

	// --- Length of map ---

	fmt.Printf("length of map %d\n", len(states))

	// --- Slice a slice ---

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("slice = %v\n", slice)
	fmt.Printf("slice[:4] = %v\n", slice[:4])
	fmt.Printf("slice[4:] = %v\n", slice[4:])
	// Note: slice[4:] is exactly the same as slice[4:len(slice)].
	fmt.Printf("slice[4:len(slice) = %v\n", slice[4:])
	fmt.Printf("slice[2:6] = %v\n", slice[2:6])
	// Note: slice[0:0] returns [], an empty slice.
	fmt.Printf("slice[0:0] = %v\n", slice[0:0])
	// Note: slice[:] returns slice, the full slice.
	fmt.Printf("slice[:] = %v\n", slice[:])

	// --- Copy ---
	smallerSlice := []int{10, 11, 12, 13, 14}
	// Copy the minimum length of the 2 slices in the arguments.
	// Copy from left to right.
	n := copy(smallerSlice, slice)
	fmt.Printf("smaller slice: %v\n", smallerSlice)
	fmt.Printf("# elements copied: %d\n", n)
	// Copy from right to left.
	n = copy(slice, smallerSlice)
	fmt.Printf("smaller slice: %v\n", smallerSlice)
	fmt.Printf("# elements copied: %d\n", n)
}
