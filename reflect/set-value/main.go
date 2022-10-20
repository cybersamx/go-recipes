package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Class struct {
	Students []string
}

// setToSlice mutates the reflect.Value `sliceRVal` depending on the underlying type of the parameter `itemRVal`.
// If `itemRVal` is a slice, we assign `itemVal` completely to `sliceRVal`. If `itemVal` matches the type of the
// underlying item types in `sliceVal`, we append `itemRVal` to `sliceRVal`.

func setToSlice(sliceRVal, itemRVal reflect.Value) (rerr error) {
	// We can pass anything to this function, which can cause the system panic. It's a
	// good idea to have a recovery from a panic.
	defer func() {
		if r := recover(); r != nil {
			rerr = fmt.Errorf("panic in append: %v", r)
		}
	}()

	// A few key conditional checks on the parameters. The above recover handler will handle edge cases.
	if sliceRVal.Kind() != reflect.Ptr {
		return errors.New("values must be addressable so that we can set a value")
	}

	// What's the value the pointer is referencing.
	sliceRVal = sliceRVal.Elem()

	if sliceRVal.Kind() != reflect.Slice {
		return errors.New("sliceRVal must be a slice")
	}

	if itemRVal.Kind() == reflect.Slice {
		sliceRVal.Set(itemRVal)
	} else {
		sliceRVal.Set(reflect.Append(sliceRVal, itemRVal))
	}

	return nil
}

func main() {
	var names []string
	sliceRVal := reflect.ValueOf(&names)

	if err := setToSlice(sliceRVal, reflect.ValueOf("Peter")); err != nil {
		fmt.Println(err)
	}
	fmt.Println(names)

	if err := setToSlice(sliceRVal, reflect.ValueOf([]string{"Stephanie", "Nancy"})); err != nil {
		fmt.Println(err)
	}
	fmt.Println(names)

	class := Class{}
	studentsRVal := reflect.ValueOf(&class).Elem().FieldByName("Students").Addr()

	if err := setToSlice(studentsRVal, reflect.ValueOf("Peter")); err != nil {
		fmt.Println(err)
	}
	fmt.Println(class)
}
