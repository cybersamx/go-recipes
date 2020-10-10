# Handy Guide to Go

A handy guide, gist, and notes for quick reference to Go. The common knowledge, tips, nuances, traps, and gotchas of the language.

## Basics

### new vs make

* `new` returns a pointer to an allocated value of type T (ie. return) and "zeros" the value.
* `make` creates slices, maps, and channels, and returns an initialized value of type of T.

Example of `new`.

```go
// The following 2 are equivalent. Assuming Person is a struct.
var john = new(Person)
john = &Person{}
```

Use `make` to create a slice. The statement is basically saying: allocate an array of 100 int values and create a slice structure with length 10 and capacity of 100, pointing to the first 10 items of the array.

```go
vals := make([]int, 10, 100)
```

Use `make` to create a map.

```go
// Create a "zero-filled" map.
rates := make(map[string]float)
```

Use `make` to create channel.

```go
// The following 2 statements are equivalent.
unbufferedChan := make(chan string)
unbufferedChan := make(chan string, 0)

// Buffered channel.
bufferedChan := make(chan string, 10)
```

### Variables

You may want to be aware of some of the gotchas in Go variables.

**Shadow Variable**

```go
x := 2

{
    // You redeclare x and now it shadows the previous variable. Hence shadow variable
	x := 4
    fmt.PrintIn(x)  // Prints 4
}

fmt.PrintIn(x)  // Prints 2
```

**Redeclaring Variables Using Short Form**

```go
x := 2
x := 3      // Compile error

// Ok now. We can redeclare as long as the variable is part of a multi-variable
// declaration where 1 of the new variable is also declared.
x, y := 3, 4
```

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

## Enum

Go doesn't have an `enum` keyword so the way we define an enum type is by using the `iota` keyword.

```go
const (
    Info  = iota    // Always start with 0, unless `iota + 1` is used to start it off with 1
    Warn
    Error
    _               // Skip a value
    Fatal
)

fmt.Println(Info, Warn, Error, Fatal)  // Prints: 0 1 2 4
```

Here's the idiomatic way for implementing a complete custom enum type in Go:

```go
// 1. Define the custom type.
type Severity int

// 2. Assign values to each enum, using to iota to increment the value.
const (
    Info = iota
    Warn
    Error
    Fatal
)

// 3. If you need the string representation of the enums, implement the String method.
func (s Severity) String() string {
    return [...]string{"Info", "Warn", "Error", "Fatal"}
}

// If Go sees that that a type implements the String() method, it will use it to print the
// enum values in string.
fmt.Println(Info, Warn, Error, Fatal)  // Prints: Info Warn Error Fatal
```


## Strings

### Immutability

A string in Go is a read-only slice of bytes, thus a string is immutable.

### #Characters vs #Bytes

## Collections

### Slice

Slice is really a struct referencing an array segment. A slice is a struct with the following fields:
  * **ptr** - to data - pointer to the data, which is an array.
  * **len** - the first len items to return from the array segment.
  * **cap** - the max length of the array segment.

### Better Way to Check Non-Existing Map Keys

In Go, when you access a map for a key that doesn't exist, Go returns a zero value for the type of the corresponding value. But checking for the appropriate zero value isn't always the right approach. For example, what if you have a map of boolean values? The returned zero value is ambiguous because it can mean that the key doesn't exist, or the value matching the key is false.

```go
// If a key isn't found in a map, Go returns the zero value for that the value type.
hasCoast := map[string]bool{"CA": true, "IL": false, "UT": false}

// This doesn't matter sense, it doesn't tell if return value isn't found or
// that it has a value of false.
if hasCoast["NY"] == false {
    fmt.Println("Not found or doesn't have coast")
}

// Here's a better way.
if val, ok := hasCoast["NY"]; ok {
    fmt.Println("Found and value is", val)
} else {
    fmt.Println("Not found")
}
```

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
