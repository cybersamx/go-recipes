package main

import (
	"fmt"
	"reflect"
)

type Loc struct {
	Lat  float64
	Long float64
}

type Person struct {
	Name string
	Age  int
	Loc
}

func main() {
	// Integer
	num := 123
	rval := reflect.ValueOf(num)
	rtyp := reflect.TypeOf(num)
	fmt.Println("--- Integer ---")
	fmt.Println(rval)        // Output: 123
	fmt.Println(rval.Kind()) // Output: int
	fmt.Println(rval.Type()) // Output: int
	fmt.Println(rtyp)        // Output: int
	//fmt.Println(rval.Elem()) // Panics since rval is references a value
	//fmt.Println(rtyp.Elem()) // Panics since rtyp is references a value
	fmt.Println()

	// Integer pointer
	rval = reflect.ValueOf(&num)
	rtyp = reflect.TypeOf(&num)
	fmt.Println("--- Integer Pointer ---")
	fmt.Println(rval)        // Output: 0xc00001a0f8 <address>
	fmt.Println(rval.Kind()) // Output: ptr
	fmt.Println(rval.Type()) // Output: *int
	fmt.Println(rtyp)        // Output: *int
	// Call method Elem() to expose the underlying value of a pointer.
	fmt.Println(rval.Elem())        // Output: 123
	fmt.Println(rval.Elem().Kind()) // Output: int
	fmt.Println(rtyp.Elem())        // Output: int

	// Slice
	texts := []string{"Bangkok", "Chicago", "Cairo"}
	rval = reflect.ValueOf(texts)
	rtyp = reflect.TypeOf(texts)
	fmt.Println("--- Slice ---")
	fmt.Println(rval)        // Output: [Bangkok Chicago Cairo]
	fmt.Println(rval.Kind()) // Output: slice
	fmt.Println(rval.Type()) // Output: []string
	fmt.Println(rtyp)        // Output: []string
	//fmt.Println(rval.Elem()) // Panics since rval is references a value
	fmt.Println(rtyp.Elem()) // Output: string
	fmt.Println()

	// Slice of Pointers
	david := "David"
	linda := "Linda"
	textPtrs := []*string{&david, &linda}
	rval = reflect.ValueOf(textPtrs)
	rtyp = reflect.TypeOf(textPtrs)
	fmt.Println("--- Slice of Pointers ---")
	fmt.Println(rval)        // Output: [0xc000014280 0xc000014290]
	fmt.Println(rval.Kind()) // Output: slice
	fmt.Println(rval.Type()) // Output: []*string
	fmt.Println(rtyp)        // Output: []*string
	//fmt.Println(rval.Elem()) // Panics since rval is references a value
	fmt.Println(rtyp.Elem()) // Output: *string
	fmt.Println()

	// Pointer to Slice
	rval = reflect.ValueOf(&texts)
	rtyp = reflect.TypeOf(&texts)
	fmt.Println("--- Pointer to a Slice of Pointers ---")
	fmt.Println(rval)        // Output: &[Bangkok Chicago Cairo]
	fmt.Println(rval.Kind()) // Output: ptr
	fmt.Println(rval.Type()) // Output: *[]string
	fmt.Println(rtyp)        // Output: *[]string
	fmt.Println(rval.Elem()) // Output: [Bangkok Chicago Cairo]
	fmt.Println(rtyp.Elem()) // Output: []string
	fmt.Println()

	// Struct
	person := Person{
		Name: "Erwin",
		Age:  20,
		Loc: Loc{
			Lat:  45.56,
			Long: -115.49,
		},
	}
	rval = reflect.ValueOf(person)
	rtyp = reflect.TypeOf(person)
	fmt.Println("--- Struct ---")
	fmt.Println(rval)        // Output: {Erwin 20 {45.56 -115.49}}
	fmt.Println(rval.Kind()) // Output: struct
	fmt.Println(rval.Type()) // Output: main.Person
	fmt.Println(rtyp)        // Output: main.Person
	//fmt.Println(rval.Elem()) // Panics since rval is references a value
	//fmt.Println(rtyp.Elem()) // Panics since rtyp is references a value
	fmt.Println()

	rval = reflect.ValueOf(&person)
	rtyp = reflect.TypeOf(&person)
	fmt.Println("--- Pointer to Struct ---")
	fmt.Println(rval)        // Output: &{Erwin 20 {45.56 -115.49}}
	fmt.Println(rval.Kind()) // Output: ptr
	fmt.Println(rval.Type()) // Output: *main.Person
	fmt.Println(rtyp)        // Output: *main.Person
	fmt.Println(rval.Elem()) // Output: {Erwin 20 {45.56 -115.49}}
	fmt.Println(rtyp.Elem()) // Output: main.Person
	fmt.Println()
}
