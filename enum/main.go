package main

import (
	"fmt"
)

// Enums in Go

type MyEnum int

const (
	MyEnumVal1 MyEnum = iota
	MyEnumVal2
	MyEnumVal3
)

var strs = [...]string{"MyEnumVal1", "MyEnumVal2", "MyEnumVal3"}

// String returns the enum values as string.
func (me MyEnum) String() string {
	return strs[me]
}

func main() {
	fmt.Println(MyEnumVal1, MyEnumVal2, MyEnumVal3)
}
