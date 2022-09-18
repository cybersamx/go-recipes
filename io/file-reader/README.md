# File Reader

An example of reading from a text file.

## Best Practices

### Pass by io.Reader and io.Writer

Go provides the `io.Reader` and `io.Writer` interfaces for I/O operations. As a best practice, pass any I/O stream as `io.Reader` or `io.Writer` instead of `string`, `[]bytes`, etc. This makes the code cleaner and more extensible.

### ReadCloser

In Go, `io.Reader` is a stream. So once the reader has finished reading a stream, you can't reread the stream again. For example:

```go
fileReader, err := os.Open("testdata.txt")
if err != nil {
	panic(err)
}
defer fileReader.Close()

data, err := io.ReadAll(fileReader)
fmt.Println(string(data)) // Prints the content of testdata.txt as string.

data, err = io.ReadAll(fileReader)
fmt.Println(string(data)) // Prints nothing.
```

To reread a stream is simply to instantiate a new instance of `io.Reader` to read the content from the stream again.

```go
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func main() {
	printFile("testdata.txt")
	printFile("testdata.txt")
}
```

### Reading in chunks

The function `io.ReadAll` reads the content from the stream into memory until an error or EOF. For bigger files, it makes more sense to break the read operations into smaller chunks.

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference and Credits

* [Go Package io](https://golang.org/pkg/io/)
