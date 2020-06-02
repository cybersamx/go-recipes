# Unicode

Unicode is a large lookup database of glyphs and symbols, which can be individually be referenced by a Unicode symbol can be referenced by a **Code Point**.

## Basic Multilingual Plane

The naming convention of a `Unicode code point` is `U+XXXX` where XXXX is zero-padded hexadecimal number ranging from `0000` to `FFFF`, which is also known as the **Basic Multilingual Plane (BMP)**. BMP code point can be accessed as a single code unit in UTF-16 encoding ie. basically a 16-bit data type to represent a Unicode character with a theoretical range of U+0000 to U+FFFF.

## Supplementary Plane

BMP was designed to be a 16-bit encoding. However, the initial set of 65,536 distinct symbols quickly ran out, and the Unicode consortium extended the Unicode set by another 1+ million symbols with the addition of 2 digits. Known as the **Supplementary Plane (SP)**, the characters in the extended set are in the range U+010000 to U+10FFFF. Read [here](https://en.wikipedia.org/wiki/Unicode) for more information on Unicode.

For an SP the code point is 21-bit, and by using a UTF-32 to encode a Unicode character would seem space inefficient as the 11 bits are not being used. Most programming language uses UTF-16 encoding to represent a supplementary character to be space efficient. So a 21-bit supplementary character is represented by a pair of surrogate 16-bit code points in a UTF-16 encoded string environment. For instance the symbol 👍 is U+1F44D in SP, but it can also be represented as U+D83D and U+DC4D (for strings encoded in UTF-16).

> **Note**
> As you can see, any number in the BMP code point in the range of U+0000 to U+D7FF and U+D7FF and U+E000 to U+10FFFF represents a Unicode character. A code point in range U+D800 to U+D7FF represents a surrogate part code point for a supplementary character.

## Ranges

| Name                      | Lower range | Upper range |
|---------------------------|-------------|-------------|
| Basic Multilingual Plane  | U+0000      | U+D7FF      |
| Surrogate Pair Code Point | U+D800      | U+DFFF      |
| Supplementary Plane       | U+100000    | U+10FFFF    |

## Extended Grapheme Clusters

An **Extended Grapheme Cluster** is a sequence of 1 or more Unicode character that are combined to produce a single human-readable character. You can use it to represent characters in decomposed or precomposed forms. For example:

```go
fmt.Println("\u00E1")         // Prints á using a precomposed á
fmt.Println("\u0061\u0301")   // Prints á using a decomposed combo of a, ◌́
```

* [Javascript has a unicode problem](https://mathiasbynens.be/notes/javascript-unicode)
* [Wikipedia: Unicode](https://en.wikipedia.org/wiki/Unicode)
* [Unicode code converter](https://r12a.github.io/app-conversion/)
* [Swift strings and characters](https://docs.swift.org/swift-book/LanguageGuide/StringsAndCharacters.html)
