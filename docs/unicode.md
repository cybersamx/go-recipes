# Unicode

Unicode is a large lookup database of glyphs and symbols, which can be individually be referenced by a Unicode symbol can be referenced by a **Code Point**.

## Basic Multilingual Plane

The naming convention of a `Unicode code point` is `U+XXXX` where XXXX is zero-padded hexadecimal number ranging from `0000` to `FFFF`, which is also known as the **Basic Multilingual Plane (BMP)**. BMP code point can be accessed as a single code unit in UTF-16 encoding ie. basically a 16-bit data type to represent a Unicode character with a theoretical range of U+0000 to U+FFFF.

## Supplementary Plane

BMP was designed to be a 16-bit encoding. However, the initial set of 65,536 distinct symbols quickly ran out, and the Unicode consortium extended the Unicode set by another 1+ million symbols with the addition of 2 digits. Known as the **Supplementary Plane (SP)**, the characters in the extended set are in the range U+010000 to U+10FFFF. Read [here](https://en.wikipedia.org/wiki/Unicode) for more information on Unicode.

For an SP the code point is 21-bit, and by using a UTF-32 to encode a Unicode character would seem space inefficient as the 11 bits are not being used. Most programming language uses UTF-16 encoding to represent a supplementary character to be space efficient. So a 21-bit supplementary character is represented by a pair of surrogate 16-bit code points in a UTF-16 encoded string environment. For instance the symbol 👍 is U+1F44D in SP, but it can also be represented as U+D83D and U+DC4D (for strings encoded in UTF-16).

> **Note**
>
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

## Encoding

Unicode provides a clean abstraction and standards that allows character to be represented in computer system independent of the programming language. **Unicode doesn't dictate how the characters should be represented in computer systems such data storage and data transmission.**

Strings in computer systems are just an arbitrary length of bytes. So we need a sequence of bit patterns or numbers to construct a sequence of code points like Unicode or ASCII characters, which can then be finally be recomposed in a human readable text.

The sequence of bit pattens in computer system is called [encoding](https://en.wikipedia.org/wiki/Character_encoding). A string in a computer system is meaningless unless we know its encoding. Indeed, "there ain't no such thing as plain text," Joel Spolsky, [Joel on Software](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/). For the context of this document, we will focus on 2 popular encodings: [ASCII]() and [UTF-8](https://en.wikipedia.org/wiki/UTF-8). Note that UTF-8 is backwards compatibility with ASCII and accounts for [about 95% of all web pages as of 2020](https://w3techs.com/technologies/cross/character_encoding/ranking). So in some ways, you can still assume an ASCII encoding for more plain text.

### ASCII

Characters (mostly based on the English alphabet) in early computer systems are represented as a 7-bit number. So the following text is represented a sequence of 7-bit hexadecimals in ASCII encoding:

|    h   |    e   |    l   |    l   |    o   |
|-------:|-------:|-------:|-------:|-------:|
| 0x68   | 0x65   | 0x6C   | 0x6C   | 0x6F   |

> **Notes**
>
> A single hexadecimal is 4-bit long. A byte is represented by 2 hexadecimals.

ASCII is a widely accepted encoding standards as it's space-efficient and simple to understand and implemented. Its biggest constraint is that it is only 7-bit wide and hence limited to 128 characters - not enough to accommodate for characters in other languages.

To enable backward compatibility with ASCII, the first 128 characters of Unicode are mapped to the first 128 characters of ASCII.

### UTF-8

8-bit Unicode Transformation Format or UTF-8 is designed to encode Unicode characters into a text string while maintaining backward compatibility with ASCII.

Also, UTF-8 is a variable width encoding. It can encode a Unicode code point in 1, 2, 3, or 4 bytes depending on the significant bits in the numerical value of the code point. Basically if the eighth bit is 0, decode the current byte as ASCII. Otherwise (if the eighth bit is 1), parse the next 4 bits to determine the size of the unicode code point and then parse according. Here's the breakdown:

| # Bytes | Lower code point | Upper code point | Byte 1    | Byte 2    | Byte 3    | Byte 4    |
|:-------:|-----------------:|-----------------:|----------:|----------:|----------:|----------:|
| 1       | U+0000           | U+007F           | 0xxx xxxx |           |           |           |
| 2       | U+0080           | U+07FF           | 110x xxxx | 10xx xxxx |           |           |
| 3       | U+0800           | U+FFFF           | 1110 xxxx | 10xx xxxx | 10xx xxxx |           |
| 4       | U+10000          | U+10FFFF         | 1111 0xxx | 10xx xxxx | 10xx xxxx | 10xx xxxx |

So with the same text "hello" has the same byte ordering in UTF-8 encoding as it does in ASCII encoding:

|    h   |    e   |    l   |    l   |    o   |
|-------:|-------:|-------:|-------:|-------:|
| U+0068 | U+0065 | U+006C | U+006C | U+006F |
| 0x68   | 0x65   | 0x6C   | 0x6C   | 0x6F   |

And text "你好世界" with 3-byte unicode in UTF-8 encoding looks like this:

|    你     |     好    |     世    |     界    |
|---------:|---------:|---------:|---------:|
| U+4F60   | U+597D   | U+4E16   | U+754C   |
| 0xE4BDA0 | 0xE5A5BD | 0xE4B896 | 0xE7958C |

## Reference

* [Javascript has a unicode problem](https://mathiasbynens.be/notes/javascript-unicode)
* [Wikipedia: Unicode](https://en.wikipedia.org/wiki/Unicode)
* [Wikipedia: UTF-8](https://en.wikipedia.org/wiki/UTF-8)
* [Unicode code converter](https://r12a.github.io/app-conversion/)
* [Swift strings and characters](https://docs.swift.org/swift-book/LanguageGuide/StringsAndCharacters.html)
