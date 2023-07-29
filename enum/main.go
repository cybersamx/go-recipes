package main

import (
	"errors"
	"fmt"
)

// Enums in Go. Explicitly define the String() method to convert a enum value to a string value.

type MyEnum int

const (
	MyEnumVal1 MyEnum = iota
	MyEnumVal2
	MyEnumVal3
)

var errMyEnumParse = errors.New("can't parse enum")

var strs = [...]string{"MyEnumVal1", "MyEnumVal2", "MyEnumVal3"}

// String returns the enum values as string.
func (me MyEnum) String() string {
	return strs[me]
}

func ParseMyEnum(text string) (MyEnum, error) {
	for i, str := range strs {
		if str == text {
			return MyEnum(i), nil
		}
	}

	return MyEnum(0), errMyEnumParse
}

// The implicit way by using go:generate.

type MyEnumImplicit int

//go:generate stringer -type=MyEnumImplicit
const (
	MyEnumValImplicit1 MyEnumImplicit = iota
	MyEnumValImplicit2
	MyEnumValImplicit3
)

func main() {
	fmt.Println(MyEnumVal1, MyEnumVal2, MyEnumVal3)

	texts := []string{"MyEnumVal1", "MyEnumVal2", "MyEnumVal3", "NonExistEnum"}
	for _, text := range texts {
		val, err := ParseMyEnum(text)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(val)
	}

	fmt.Println(MyEnumValImplicit1, MyEnumValImplicit2, MyEnumValImplicit3)
}
