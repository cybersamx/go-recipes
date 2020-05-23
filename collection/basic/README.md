# Collections in Go

This recipe highlights common operations on slice and map as well as some relevant tips and pitfalls.

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

### You can append to a nil or empty (or initialized) slice

Go `append` function works for both nil and empty slices. Here's an example for nil slice.

```go
var nilSlice []State

// Append to a nil slice.
nilSlice = append(nilSlice, State{name: "California", population: 39512223})
fmt.Println(nilSlice)
```

Works the same for an empty slice.

```go
var emptySlce = []State{}

// Append to an empty slice.
emptySlce = append(emptySlce, State{name: "California", population: 39512223})
fmt.Println(emptySlce)
```

> **Note**
>
> We can think of the `append` function to be a 2-step operation that create a new slice and copy content form one slice to the new slice.
> ```go
> b := append([]T{}, a)
> ```
> 
> is the same as:
>
> ```go
> b := make([]T{}, len(a))
> copy(b, a)
> ```

### You can only add to an empty (or initialized)

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

### Subslice

In this [blog](https://blog.golang.org/slices), Rob Pike describes it best:

"A slice is a data structure describing a contiguous section of an array stored separately from the slice variable itself. _**A slice is not an array. A slice describes a piece of an array.**_"

In other words, we can create slice out of a slice. And the new subslice isn't a copy of the original slice, it's merely a slice of the original slace - all data stays. 

We describes the boundaries of subslice by using the `[:]` operator. Let's use the following slice as a reference:

```go
s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
```

| Operator | Output                | Description                                       |
|----------|-----------------------|---------------------------------------------------|
| s[:4]    | [0 1 2 3]             | slice with elements from start thru index (4-1)   |
| s[4:]    | [4 5 6 7 8 9]         | slice with elements from index 4 thru end         |
| s[2:6]   | [2 3 4 5]             | slice with elements from index 2 thru index (6-1) |
| s[0:0]   | []                    | slice with 0 elements, an empty slice             |
| s[:]     | [0 1 2 3 4 5 6 7 8 9] | full slice inclusive of all elements              |

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference

* [The Go Blog: Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices)
* [The Go Blog: Go maps in action](https://blog.golang.org/maps)
