# Go Map

This recipe highlights common operations as well as some relevant tips and pitfalls on using map in Go.

## Tips and Gotchas

### Getting non-existing value from map

In Go, a map returns a zero-value if key isn't found through the `[key]` operator. And zero value is contextual depending on the data type of the map values. If the map value is declared as a pointer, the map returns a `nil`, but if the map value is declared as a value, the map returns a zero value corresponding to the type. Because of this, when checking for the existence of a key, use the second parameter that's returned from the `[key]` operator.

Here we declare the values of our map as a value type:

```go
	type State struct {
		name string
		population int
	}

	states := map[string]State{
		"CA": State{name: "California", population: 39512223},
		"MA": State{name: "Massachusetts", population: 6892503},
	}

	state := states["WA"]
	fmt.Println(state)  // Prints {"", "", 0}

	// Better
	state, ok := states["WA"]
	fmt.Println("Does it exist? ", ok)  // Prints false
```

Here we declare the values of our map as pointers:

```go
	states := map[string]*State{
		"CA": &State{name: "California", population: 39512223},
		"MA": &State{name: "Massachusetts", population: 6892503},
	}

	state := states["WA"]
	fmt.Println(state)  // Prints <nil>
```

### You can only add to an empty (or initialized) map

The following will cause `panic: assignment to entry in nil map` at runtime.

```go
var states map[string]State
states["CA"] = State{name: "California", population: 39512223}  // Panics
```

We need to initialize the map before we can add data to it.

```go
states := make(map[string]State)
// This works too: var states = map[string]State{}
states["CA"] = State{name: "California", population: 39512223}  // Works now
```

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference

* [The Go Blog: Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices)
* [The Go Blog: Go maps in action](https://blog.golang.org/maps)
