# Count and Iterate Over a Go String

A simple recipe that explores how we count and iterate over a string by rune and byte.

Unicode can be complex, but it isn't rocket science. Here's a short [primer](../../docs/unicode.md) to understand the basics of Unicode and text encoding. Also, read the [blog on strings, bytes, and runes, and characters](https://blog.golang.org/strings) on the official Go website.

## Definitions

| Term           | Description  |
|----------------|--------------|
| String         | String is a read-only slice of arbitrary bytes encoded in UTF-8 (absent of byte-level escapes. |
| Code point     | A numerical value that is mapped to a character. It's dependent on the character encoding eg. ASCII or Unicode. |
| Rune           | Go speak for code point. |
| Byte           | 8-bit or 1-byte of a unit of digital information. |
| Character      | An abstract representation of a symbol. The term character is often ambiguous and it really depends on the context given a character can be represented in different ways. |

## Byte Count vs Rune Count

Let's use this string as an example: 你好世界. In Go, we can represent it in the following ways:

```go
// Both strings are equivalent.
str := "你好世界"   // As UTF-8 string.
str := "\xe4\xbd\xa0\xe5\xa5\xbd\xe4\xb8\x96\xe7\x95\x8c"   // As bytes using byte-level escapes.
```

How big is a string? You can count the size by number of bytes or by number of runes.

```go
fmt.Println(len(str))                       // Prints 12
fmt.Println(utf8.RuneCountInString(str))    // Prints 4
```

See [source code](main.go) for details.

## Iterate by Byte vs Iterate by Byte

To iterate over a string by byte, here are the general techniques:

```go
for i := 0; i < len(str); i++ {
	fmt.Printf("%x ", str[i])
}

for _, b := range []byte(str) {
	fmt.Printf("%x ", b)
}
```

To iterate over a string by rune, here are the general techniques:

```go
for i, b := range str {
	fmt.Printf("%q starts at position (in byte) %d ", b, i)
}
```

Pay close attention to the index `i`, it denotes the index of the rune in byte units.

See [source code](main.go) for details.

## Setup

1. Run the program

   ```bash
   $ make
   ```

## Credits and Reference

* [The Go Blog: Strings, bytes, and runes, and characters](https://blog.golang.org/strings)
* [The Go Blog: Text normalization in Go](https://blog.golang.org/normalization)
* [Local Unicode Primer](../../docs/unicode.md)
* [Joel On Software: The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](http://www.joelonsoftware.com/articles/Unicode.html)
