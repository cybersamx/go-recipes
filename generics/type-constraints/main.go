package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func double[T int | int32 | int64 | float32 | float64](t []T) {
	for i, v := range t {
		t[i] = v + v
	}
}

// We can define a type constraint, a new Go construct to represent all the supported
// types.

type Num interface {
	int | int32 | int64 | float32 | float64
}

// doubleNum declares T as Num instead of writing out int | int32 | int64...
func doubleNum[T Num](t []T) {
	for i, v := range t {
		t[i] = v + v
	}
}

// Num is good but not great. What if we have a type definition like MyInteger. We can
// pass an MyInteger value to doubleNum even though Integer is derived from int.

type MyInteger int

// Use type approximation by prefixing the base type with ~.

type ImprovedNum interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// doubleImprovedNum uses type approximation.
func doubleImprovedNum[T ImprovedNum](t []T) {
	for i, v := range t {
		t[i] = v + v
	}
}

// doubleOrdered uses the Ordered type in the type-constraints package.
func doubleOrdered[T constraints.Ordered](t []T) {
	for i, v := range t {
		t[i] = v + v
	}
}

// Nothing happens at runtime because the functions below aren't called. It's here to
// illustrate what types of operators that are supported for a given type constraint. We
// can determine this at compile time. If you uncomment the one of the comments with an
// operator statement, this program won't build.

func testAny[T any](a, b T) {
	// These operators are not defined for T.
	//_ = a == b
	// _ = a != b
	// _ = a >= b
	// _ = a > b
	// _ = a < b
	// _ = a <= b
	// _ = a + b
}

func testComparable[T comparable](a, b T) {
	// These operators are defined for T.
	_ = a == b
	_ = a != b

	// These operators are not defined for T.
	// _ = a >= b
	// _ = a > b
	// _ = a < b
	// _ = a <= b
	// _ = a + b
}

func testOrdered[T constraints.Ordered](a, b T) {
	// These operators are defined for T.
	_ = a == b
	_ = a != b
	_ = a >= b
	_ = a > b
	_ = a < b
	_ = a <= b
	_ = a + b
}

func main() {
	floats := []float32{1.0, 2.0, 3.0}
	ints := []int{1, 2, 3}

	double(floats)
	double(ints)
	fmt.Println(floats)
	fmt.Println(ints)

	doubleNum(floats)
	doubleNum(ints)
	fmt.Println(floats)
	fmt.Println(ints)

	integers := []MyInteger{1, 2, 3}

	// The following won't work.
	// doubleNum(integers)
	// This will work because we are using type approximation.
	doubleImprovedNum(integers)
	fmt.Println(integers)

	// Or just use the type-constraints defined in the type-constraints package.
	doubleOrdered(integers)
	fmt.Println(integers)
}
