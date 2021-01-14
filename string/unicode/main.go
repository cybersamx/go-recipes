package main

import "fmt"

func main() {
	// Hex representation
	fmt.Println("a" == "\x61") // True

	// Unicode representation
	// Use escape \u followed by 4 digit unicode code point.
	// Use escape \U followed by 8 digit unicode supplementary code point
	fmt.Println("a" == "\u0061") // True

	// Go supports supplementary Unicode code point directly.
	// Surrogate pair code points is unsupported.
	// ie. print("👍" == "\u{D83D}\u{DC4D}") doesn't work.
	fmt.Println("👍" == "\U0001F44D") // True

	// Extended grapheme cluster
	fmt.Println("\u00E1")       // Prints á using a precomposed á
	fmt.Println("\u0061\u0301") // Prints á using a decomposed combo of a, ◌́

	fmt.Println("\uD55C")             // Prints 한
	fmt.Println("\u1112\u1161\u11AB") // Prints ᄒ, ᅡ, ᆫ
}
