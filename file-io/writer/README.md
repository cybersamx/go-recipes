# File Reader

An example of writing a text file.

## Best Practices

### Pass by io.Reader and io.Writer

Go provides the `io.Reader` and `io.Writer` interfaces to represent I/O operations. As a best practice, pass the I/O streams using `io.Reader` or `io.Writer` instead of `string`, `[]bytes`, etc. As interfaces, you can pass the concrete object representing the actual stream - it makes the code cleaner and more extensible.

### Writing in chunks

The function `ioutil.ReadAll` reads the content from the stream into memory until an error or EOF. For big files, it makes more sense to break the read operations into smaller chunks.

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference and Credits

* [Go Package io](https://golang.org/pkg/io/)
* [Go Package ioutil](https://golang.org/pkg/io/ioutil/)
