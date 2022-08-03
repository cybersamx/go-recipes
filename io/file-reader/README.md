# File Reader

An example of reading from a text file.

## Best Practices

### Pass by io.Reader and io.Writer

Go provides the `io.Reader` and `io.Writer` interfaces for I/O operations. As a best practice, pass any I/O stream as `io.Reader` or `io.Writer` instead of `string`, `[]bytes`, etc. This makes the code cleaner and more extensible.

### ReadCloser

In Go, `io.Reader` is a stream. So once the reader has finishing reading a stream, you can't reread the stream again. For example:

```go
file, err := os.Open("testdata.txt")
file.Close()
data, err := ioutil.ReadAll(file)
fmt.Println(data)   // Prints the content of testdata.txt

data2, err := ioutil.ReadAll(file)
fmt.Println(data2)   // Prints nothing
```

To print the file the second time, you have to create a new stream called `io.ReadCloser` from the content that is loaded from reading the `io.Reader` stream.

```go
file, err := os.Open("testdata.txt")
file.Close()
data, err := ioutil.ReadAll(file)
fmt.Println(data)   // Prints the content of testdata.txt

file2 := ioutil.NopCloser(bytes.NewReader(data))
data2, err := ioutil.ReadAll(file2)
fmt.Println(data2)   // Prints the content of testdata.txt
```

### Reading in chunks

The function `ioutil.ReadAll` reads the content from the stream into memory until an error or EOF. For big files, it makes more sense to break the read operations into smaller chunks.

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference and Credits

* [Go Package io](https://golang.org/pkg/io/)
* [Go Package ioutil](https://golang.org/pkg/io/ioutil/)
