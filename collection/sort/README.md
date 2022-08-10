# Sort a Slice

Say we need to sort a slice of custom type elements. We can do this in 2 ways:

* We wrap the underlying slice with a custom type that implements the `sort.Interface` with the methods that
  encapsulate the sorting algorithm.
* We define a closure encapsulating the sorting algorithm and pass it to the `sort.Slice` function.

See the example for details.

## Setup

1. Run the program.

   ```bash
   $ make run
   ```

## Reference

* [The sort package](https://pkg.go.dev/sort)
