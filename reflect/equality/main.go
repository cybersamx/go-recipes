package main

import (
	"fmt"
	"reflect"
)

type State int

const (
	CA = iota
	MA
	NY
	UT
)

func (s State) String() string {
	return [...]string{"CA", "MA", "NY", "UT"}[s]
}

type Gender int

const (
	Female = iota
	Male
)

func (g Gender) String() string {
	return [...]string{"Male", "Female"}[g]
}

type Person struct {
	Name   string
	Gender Gender
	Age    int
}

type Address struct {
	StreetAddress string
	City          string
	ZipCode       string
	State         State
	Latitude      float64
	Longitude     float64
	IsResidential bool
	Occupants     []Person
	rating        int
}

func main() {
	addressA := Address{
		StreetAddress: "1111 S Figueroa St",
		City:          "Los Angeles",
		ZipCode:       "90015",
		State:         CA,
		Latitude:      34.0430175,
		Longitude:     -118.2672541,
		IsResidential: false,
		rating:        90,
		Occupants:     []Person{
			{Name: "Shaq", Gender: Male, Age: 48},
			{Name: "Kobe", Gender: Male, Age: 41},
		},
	}

	lakers := Address{
		StreetAddress: "1111 S Figueroa St",
		City:          "Los Angeles",
		ZipCode:       "90015",
		State:         CA,
		Latitude:      34.0430175,
		Longitude:     -118.2672541,
		IsResidential: false,
		rating:        90,
		Occupants:     []Person{
			{Name: "Shaq", Gender: Male, Age: 48},
			{Name: "Kobe", Gender: Male, Age: 41},
		},
	}

	reversedRoster := Address{
		StreetAddress: "1111 S Figueroa St",
		City:          "Los Angeles",
		ZipCode:       "90015",
		State:         CA,
		Latitude:      34.0430175,
		Longitude:     -118.2672541,
		IsResidential: false,
		Occupants:     []Person{
			{Name: "Kobe", Gender: Male, Age: 41},
			{Name: "Shaq", Gender: Male, Age: 48},
		},
	}

	addressB := Address{
		StreetAddress: "77 Massachusetts Ave",
		City:          "Cambridge",
		ZipCode:       "02139",
		State:         MA,
		Latitude:      42.407211,
		Longitude:     -71.382439,
		IsResidential: false,
		Occupants:     nil,
	}

	// Different values of the same type.
	ok := reflect.DeepEqual(addressA, addressB)
	fmt.Println("Are addressA and addressB equal?", ok)

	// Same values of the same type.
	ok = reflect.DeepEqual(addressA, lakers)
	fmt.Println("Are addressA and lakers equal?", ok)
	fmt.Printf("Address of addressA: %p\n", &addressA)
	fmt.Printf("Address of lakers: %p\n", &lakers)

	// Same values but ordering changed.
	ok = reflect.DeepEqual(addressA, reversedRoster)
	fmt.Println("Are addressA and reversedRoster equal?", ok)
}
