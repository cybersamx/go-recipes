package main

import "fmt"

// Using a custom type.
type State struct {
	name         string
	abbreviation string
	population   int
}

func main() {
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
		name:         "California",
		abbreviation: "CA",
		population:   39512223,
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
	fmt.Printf("does CA exist? %s %t\n", state, ok)

	// --- Iterate the map ---

	ma := State{
		name:         "Massachusetts",
		abbreviation: "MA",
		population:   6892503,
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
