# File Reader

An example of reading from a text file.

## ReadAll vs Read

The function `io.ReadAll` reads the content from the stream into memory until an error or EOF. For bigger files, it makes more sense to break the read operations into smaller chunks.

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference and Credits

* [Go Package io](https://golang.org/pkg/io/)
