# Multi-Read

Reading a source stream multiple times.

## Overview

In Go, `io.Reader` is a stream. So once the reader has finished reading a stream, you can't reread the stream again. For example:

```go
str := "hello world"
reader := bytes.NewBufferString(str)
buf := make([]byte, len(str))

reader.Read(buf)
fmt.Println(string(buf))

// Reading it again returns nothing because the stream has been read.
rebuf := make([]byte, len(str))
reader.Read(rebuf)
fmt.Println(string(rebuf))
```

To reread a stream we instantiate a new `io.Reader` so that we can read the content from the stream again.

## Reference and Credits

* [Go Package io](https://golang.org/pkg/io/)
