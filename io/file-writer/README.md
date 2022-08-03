# File Reader

An example of writing to a text file.

## Best Practices

### Pass by io.Reader and io.Writer

Go provides the `io.Reader` and `io.Writer` interfaces for I/O operations. As a best practice, pass any I/O stream as `io.Reader` or `io.Writer` instead of `string`, `[]bytes`, etc. This makes the code cleaner and more extensible.

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference and Credits

* [Go Package io](https://golang.org/pkg/io/)
* [Go Package ioutil](https://golang.org/pkg/io/ioutil/)
