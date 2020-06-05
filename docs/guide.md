# Handy Guide to Go

A handy guide, gist, and notes for quick reference to Go. The common knowledge, tips, nuances, traps, and gotchas of the language.

## I/O

### What Implements io.Reader and io.Writer

The `io` package in the standard Go library defines 2 interfaces `io.Reader` and `io.Writer` that represent reading from a data source and writing to a data target respectively.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Here are the structs that implement the 2 interfaces.

| Struct, Variables, Constants                                 | Reader | Writer | Description                                                                                            |
|--------------------------------------------------------------|--------|--------|--------------------------------------------------------------------------------------------------------|
| [bytes.Reader](https://golang.org/pkg/bytes/#Reader)         |   ✓    |        | Created by calling `bytes.NewReader` that returns a new Reader reading from a byte slice.              |
| [bytes.Buffer](https://golang.org/pkg/bytes/#Buffer)         |   ✓    |    ✓   | Created by declaring a `bytes.Buffer` struct, or calling `bytes.NewBuffer` or `bytes.NewBufferString`. |
| [strings.Reader](https://golang.org/pkg/strings/#Reader)     |   ✓    |        | Created by calling `strings.NewReader`, which returns a new Reader reading from a string.              |
| [strings.Builder](https://golang.org/pkg/strings/#Builder)   |        |    ✓   | Created by calling `strings.NewBuilder`, which returns a new Reader reading from a string.             |
| [bufio.Reader](https://golang.org/pkg/bufio/#Reader)         |   ✓    |        | Created by calling `bufio.NewReader` with an underlying reader.                                        |
| [bufio.Writer](https://golang.org/pkg/bufio/#Writer)         |        |    ✓   | Created by calling `bufio.NewWriter` with an underlying writer.                                        |
| [io.LimitedReader](https://golang.org/pkg/io/#LimitedReader) |   ✓    |        | Created by declaring a `io.LimitedReader` with an underlying reader.                                   |
| [io.PipeReader](https://golang.org/pkg/io/#PipeReader)       |   ✓    |        | Created by declaring a `io.PipeReader` struct.                                                         |
| [io.PipeWriter](https://golang.org/pkg/io/#PipeWriter)       |        |    ✓   | Created by declaring a `io.PipeWriter` struct.                                                         |
| [io.SectionReader](https://golang.org/pkg/io/#SectionReader) |   ✓    |        | Created by calling `io.NewSectionReader` with an underlying reader.                                    |
| [os.File](https://golang.org/pkg/os/#File)                   |   ✓    |    ✓   | Create by calling `os.Create`, `os.NewFile`, `os.Open`, or `os.OpenFile`.                              |
| os.Stdout, os.Stdin, os.Stderr                               |   ✓    |    ✓   | These are public variables exposed by the `os` package. They are special files of type `os.File`.      |

## Reference

* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)
* [The Go Blog](https://blog.golang.org/)
