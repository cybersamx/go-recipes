# Handy Guide to Go

A handy guide, gist, and notes for quick reference to Go. The common knowledge, tips, nuances, traps, and gotchas of the language.

## Basics

### Loops

Here are the different ways to iterate through a slice in Go:

```go
items := []int{1, 2, 3}
size := len(items)

// The for-each loop pattern using for range
for i, item := range items {
	fmt.Printf("%d: %d\n", i, item)
}

// The while loop pattern using for expression
i := 0
for i < size {
	fmt.Printf("%d: %d\n", i, items[i])
	i++
}

// The for ;; loop pattern
for i := 0; i < size; i++ {
	fmt.Printf("%d: %d\n", i, items[i])
}
```

For map, we do the following:

```go
rates := map[string]float64 {"EUR": 1.0, "USD": 1.1305, "JPY": 126.40}

for k, v := range rates {
	fmt.Printf("%s: %.2f\n", k, v)
}
```

### Equality

Equality in programming languages can be generally defined as

* **Physical equality** - if two references reference the same object.
* **Structural equality** - if the two objects have the same representation ie. same contents.

The following table is taken from [Wikipedia: Relational operator](https://en.wikipedia.org/wiki/Relational_operator), which summarizes the different ways of testing for physical and structural equality in various programming languages.

| Language    | Structural Equality      | Physical Equality     |
|-------------|--------------------------|-----------------------|
| Go          | reflect.DeepEqual(a, b)  | &a == &b              |
|             | a == b                   |                       |
| Java        | a.equals(b)              | a == b                |
| Javascript  | a == b                   | a === b               |
| Python      | a == b                   | a is b                |
| Ruby        | a == b                   | a.equal?(b)           |
| Swift       | a == b                   | a === b               |

When comparing operands of the same type, the result is straightforward. In dynamic languages, there's room for interpretation when comparing operands of different types eg. integer with string. Given that Go is a strongly typed language, its relational operators compare the two operands only if their underlying types are the same type with each other. There are still some nuances but compared to Javascript, Python, and Ruby, the exceptions are far fewer.

Here's an equality table for Go.

![Equality table for Go == operator](images/equality_table.png)

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
