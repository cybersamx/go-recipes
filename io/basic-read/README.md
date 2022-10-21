# Basics

Many ways to read.

## The Read Method

Go provides the `io.Reader` interface for input operations.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

When we call `Read`, the method expects a `[]byte` argument. So we pass it a literal byte slice but the slice needs
to be of right size so that we can read and hold the content.

```go
str := "hello wold"
reader := bytes.NewBufferString(str)
buf := make([]byte, len(str))
reader.Read(buf)
fmt.Println(string(buf))
```

There are cases where a variable size buffer that grows as needed is useful. For that, we use `bytes.Buffer` as a
utility for reading streams.

```go
reader = bytes.NewBufferString(str)
buffer := new(bytes.Buffer)
buffer.ReadFrom(reader)
fmt.Println(buffer.String())
```

## Pass by io.Reader and io.Writer

Go provides the `io.Reader` and `io.Writer` interfaces for I/O operations. As a best practice, pass any I/O stream as `io.Reader` or `io.Writer` instead of `string`, `[]bytes`, etc. This makes the code cleaner and more extensible.

## Reference and Credits

* [Go Package io](https://golang.org/pkg/io/)
