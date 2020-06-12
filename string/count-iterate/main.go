package main

import (
	"fmt"
	"unicode/utf8"
)

func iterateByByte(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("0x%x ", str[i])
	}
	fmt.Println()
}

func iterateByByteUsingRange(str string) {
	for _, b := range []byte(str) {
		fmt.Printf("0x%x ", b)
	}
	fmt.Println()
}

func iterateByRuneUsingRange(str string) {
	for i, b := range str {
		// Pay close attention to the index i - it denotes the index in bytes.
		fmt.Printf("%q starts at position (in byte) %d ", b, i)
	}
	fmt.Println()
}

func iterateUsingDecodeRuneInString(str string) {
	for i, w := 0, 0; i < len(str); i += w {
		token := str[i:]
		// DecodeRune extracts the first rune in the passed string.
		runeValue, width := utf8.DecodeRuneInString(token)
		fmt.Printf("%q starts at position (in byte) %d ", runeValue, i)
		w = width
	}
	fmt.Println()
}

func main() {
	ascii := "hello world"
	unicode := "你好世界"
	suppUnicode := "\U0001F137\U0001F134\U0001F13B\U0001F13B\U0001F13E \U0001F146\U0001F13E\U0001F141\U0001F13B\U0001F133"
	mixed := "\U0001F137ell\U0001F13E 世界"

	// Print the strings.
	fmt.Println("Print the strings")
	fmt.Println(ascii)
	fmt.Println(unicode)
	fmt.Println(suppUnicode)
	fmt.Println(mixed)
	fmt.Println()

	// Number of bytes.
	fmt.Println("Bytes count")
	fmt.Println("#bytes ascii:", len(ascii))
	fmt.Println("#bytes unicode:", len(unicode))
	fmt.Println("#bytes suppUnicode:", len(suppUnicode))
	fmt.Println("#bytes mixed:", len(mixed))
	fmt.Println()

	// Number of characters (runes).
	fmt.Println("Rune count")
	fmt.Println("#runes ascii:", utf8.RuneCountInString(ascii))
	fmt.Println("#runes unicode:", utf8.RuneCountInString(unicode))
	fmt.Println("#runes suppUnicode:", utf8.RuneCountInString(suppUnicode))
	fmt.Println("#runes mixed:", utf8.RuneCountInString(mixed))
	fmt.Println()

	// Iterate by byte.
	fmt.Println("Iterate by byte")
	iterateByByte(ascii)
	iterateByByte(unicode)
	iterateByByte(suppUnicode)
	iterateByByte(mixed)

	iterateByByteUsingRange(ascii)
	iterateByByteUsingRange(unicode)
	iterateByByteUsingRange(suppUnicode)
	iterateByByteUsingRange(mixed)
	fmt.Println()

	// Iterate by rune.
	fmt.Println("Iterate by rune")
	iterateByRuneUsingRange(ascii)
	iterateByRuneUsingRange(unicode)
	iterateByRuneUsingRange(suppUnicode)
	iterateByRuneUsingRange(mixed)
	fmt.Println()

	// Get a rune from a mixed string.
	// For this example, we'll use the string with unicode characters with varying widths.
	fmt.Println("Iterate using DecodeRuneInString")
	iterateUsingDecodeRuneInString(mixed)
	fmt.Println()
}
