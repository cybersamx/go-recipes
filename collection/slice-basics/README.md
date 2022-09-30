# Go Slice

This example highlights common operations, including tips and pitfalls, on Go slice.

## Tips and Gotchas

### You can append to a nil or empty (or initialized) slice

Go `append` function works for both nil and empty slices. Here's an example for nil slice.

```go
var nilslice []State

// Append to a nil slice.
nilslice = append(nilslice, State{name: "California", population: 39512223})
fmt.Println(nilslice)
```

Works with an empty slice.

```go
var emptyslice = []State{}

// Append to an empty slice.
emptyslice = append(emptyslice, State{name: "California", population: 39512223})
fmt.Println(emptyslice)
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

### Subslice

In this [blog](https://blog.golang.org/slices), Rob Pike describes it best:

"A slice is a data structure describing a contiguous section of an array stored separately from the slice variable itself. _**A slice is not an array. A slice describes a piece of an array.**_"

In other words, we can create slice out of a slice. And the new subslice isn't a copy of the original slice, it's merely a slice of the original slice - all data stays.

We describe the boundaries of subslice by using the `[:]` operator. Let's use the following slice as a reference:

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

Use `[:]` to remove an element from a slice. For example:

```go
numbers := []int{1, 2, 3, 4}
numbers = numbers[1:]
fmt.Println(numbers)  // Output 2, 3, 4
```

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference

* [The Go Blog: Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices)
* [The Go Blog: Go maps in action](https://blog.golang.org/maps)
