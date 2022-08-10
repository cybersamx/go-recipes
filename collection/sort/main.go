// Based on the example on https://pkg.go.dev/sort

package main

import (
	"fmt"
	"sort"
	"strings"
)

type City struct {
	Name       string
	Population int
}

func (c City) Clone() City {
	return c
}

// Wrap the underlying model to sort.
type sortWrapper []City

// Implements the sort.Interface to allow sorting of custom types.

func (sw sortWrapper) Len() int {
	return len(sw)
}

func (sw sortWrapper) Less(i, j int) bool {
	return strings.Compare(sw[i].Name, sw[j].Name) < 0
}

func (sw sortWrapper) Swap(i, j int) {
	sw[i], sw[j] = sw[j], sw[i]
}

func main() {
	cities := []City{
		City{Name: "Beijing", Population: 20035455},
		City{Name: "Paris", Population: 11078546},
		City{Name: "Bangalore", Population: 13193035},
		City{Name: "Los Angeles", Population: 3909360},
		City{Name: "San Paulo", Population: 21846507},
		City{Name: "Nairobi", Population: 5118844},
	}

	var cities1, cities2 []City
	for _, city := range cities {
		cities1 = append(cities1, city.Clone())
		cities2 = append(cities2, city.Clone())
	}

	fmt.Println("unsorted cities1:", cities1)
	fmt.Println("unsorted cities2:", cities2)

	// We can sort by implement wrap the underlying value with methods that define
	// the pertinent methods to help with the sort.
	sort.Sort(sortWrapper(cities1))
	fmt.Println("sorted cities1:", cities1)

	// Another method is to just sort it with a closure.
	sort.Slice(cities2, func(i, j int) bool {
		return cities2[i].Name < cities[j].Name
	})
	fmt.Println("sorted cities2:", cities2)
}
