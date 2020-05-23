package main

import "fmt"

// Using a custom type.
type State struct {
	name string
	abbreviation string
	population int
}

func basicSlice() {
	fmt.Println("\n--- Slice ---")

	// Slice is nil.
	// If you want to instantiate and initialize a slice at the same time, do the following:
	// var states []state{}
	var states []State

	// For slice, initialization is optional. You can add an item to a nil slice.

	// --- Add an item ---
	ca := State{
		name: "California",
		abbreviation: "CA",
		population: 39512223,
	}
	states = append(states, ca)
	fmt.Println(states)

	// --- Get an item ---

	state := states[0]
	fmt.Println(state)

	// Panics if index is out of bound.

	// --- Iterate the map ---

	ma := State{
		name: "Massachusetts",
		abbreviation: "MA",
		population: 6892503,
	}
	states = append(states, ma)

	for i, item := range states {
		fmt.Printf("index: %d, item: %v\n", i, item)
	}

	// --- Length of map ---

	fmt.Printf("length of map %d\n", len(states))

	// --- Slice a slice ---

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("elements from start thru index (4-1) inclusive: %v\n", slice[:4])
	fmt.Printf("elements from index 4 inclusive thru end: %v\n", slice[4:])
	// Note: slice[4:] is exactly the same as slice[4:len(slice)].
	fmt.Printf("elements from index 4 inclusive thru end: %v\n", slice[4:len(slice)])
	fmt.Printf("elements from index 2 thru (5-1) inclusive: %v\n", slice[2:6])
	// Note: slice[0:0] returns [], an empty slice.
	fmt.Printf("elements from index 0 thru 0: %v\n", slice[0:0])
	// Note: slice[:] returns slice, the full slice.
	fmt.Printf("[:]: %v\n", slice[:])

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

func basicMap() {
	fmt.Println("\n--- Map ---")

	// Map isn't initialized and is nil.
	var states map[string]State

	// For map, you need to initialize it before you can add an item to it.
	// Adding an item to a nil map will cause a panic.
	// The following is an equivalent.
	// states = map[string]state{}
	states = make(map[string]State)

	// --- Add an item ---

	ca := State{
		name: "California",
		abbreviation: "CA",
		population: 39512223,
	}
	states["CA"] = ca
	fmt.Println(states)

	// --- Get an item ---

	// Basic
	state := states["CA"]
	fmt.Println(state)

	// Returns zero value ie. {"", "", 0} (if the item is a value) and nil (if the item is a pointer)
	// if an item isn't found.
	state = states["MA"]
	fmt.Println(state)

	// Test for existence. You can also replace state with _ given we are not interested in the value.
	state, ok := states["CA"]
	fmt.Printf("does CA exist? %t\n", ok)

	// --- Iterate the map ---

	ma := State{
		name: "Massachusetts",
		abbreviation: "MA",
		population: 6892503,
	}
	states["MA"] = ma

	// Order isn't guaranteed.
	for key, item := range states {
		fmt.Printf("key: %s, item: %v\n", key, item)
	}

	// --- Remove an item ---
	delete(states, "CA")
	_, ok = states["CA"]
	fmt.Printf("does CA exist? %t\n", ok)

	// --- Length of map ---

	fmt.Printf("length of map %d\n", len(states))
}

func main() {
	basicSlice()

	basicMap()
}
