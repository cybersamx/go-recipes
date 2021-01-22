package main

import (
	"fmt"
	"regexp"
)

func main() {
	// --- Simple Matching ---

	pattern1 := `^x.`
	sample1 := "xxyyyz123"
	ok, err := regexp.MatchString(pattern1, sample1)
	fmt.Println("Ok (should be true):", ok)
	fmt.Println("Err (should be nil):", err)

	// The function `MatchString` returns an error when the regexp pattern is invalid.
	invalidPattern := `[[:invalid:]]`
	_, err = regexp.MatchString(invalidPattern, sample1)
	fmt.Println("Err (should be non-nil):", err)

	// --- Regexp compilation ---

	// For more complex regexp pattern, compile it using `Compile` or `MustCompile`.
	// The difference between the functions is that the latter panics if the pattern is invalid.
	_, err = regexp.Compile(invalidPattern)
	fmt.Println("Err (should be non-nil):", err)
	// To recover from a panic properly we need to 1. Use defer to call recover(),
	// 2. The function that panics must be in a function scope so that once we recover
	// we return back to the parent and the flow continues.
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from a panic", r)
			}
		}()
		_ = regexp.MustCompile(invalidPattern) // Panics
	}()

	// --- Find ---

	pattern2 := `<([a-z]+)>`
	sample2 := "<title><123><TITLE><footer>"
	re2, err := regexp.Compile(pattern2)
	fmt.Println("FindString:", re2.FindString(sample2))
	fmt.Println("FindStringIndex:", re2.FindStringIndex(sample2))
	fmt.Println("FindStringSubmatch", re2.FindStringSubmatch(sample2))
	fmt.Println("FindStringSubmatchIndex", re2.FindStringSubmatchIndex(sample2))
	fmt.Println("FindAllString", re2.FindAllString(sample2, -1))
	fmt.Println("FindAllStringIndex", re2.FindAllStringIndex(sample2, -1))
	fmt.Println("FindAllStringSubmatch", re2.FindAllStringSubmatch(sample2, -1))
	fmt.Println("FindAllStringSubmatchIndex", re2.FindAllStringSubmatchIndex(sample2, -1))

	// --- Replace ---

	pattern3 := `//.*@`
	sample3 := "mongodb://nobody:secrets@localhost:27017/go-recipes"
	re3, _ := regexp.Compile(pattern3)
	replaced := re3.ReplaceAllString(sample3, "//*****:*****@")
	fmt.Println("New string", replaced)
}
