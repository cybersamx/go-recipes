package main

import "fmt"

type PetCommon struct {
	Age int
}

// Pet is a type interface. Note that in Go, interface names typically ends with 'er'.
type Pet interface {
	Say() string
}

// PeterConstraint is a type constraint interface using as a (Generic) constraint parameter.
// For future use.
type PeterConstraint interface {
	Dog | Cat
}

type Dog struct {
	PetCommon
}

type Cat struct {
	PetCommon
}

func (tc PetCommon) Say() string {
	return "Tsk"
}

// Say method in a Dog object should override the Say method in PetCommon.
func (d Dog) Say() string {
	return "Bark"
}

func main() {
	pets := []Pet{
		Dog{PetCommon{
			Age: 5,
		}},
		Cat{PetCommon{
			Age: 6,
		}},
	}

	for _, v := range pets {
		fmt.Println(v.Say())
	}
}
