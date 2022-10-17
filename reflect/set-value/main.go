package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Class struct {
	Students []string
}

// setToSlice mutates the reflect.Value `sliceVal` depending on the underlying type of the parameter `itemVal`.
// If `itemVal` is a slice, we assign `itemVal` completely to `sliceVal`. If `itemVal` matches the type of the
// underlying item types in `sliceVal`, we append `itemVal` to `sliceVal`.

func setToSlice(sliceVal, itemVal reflect.Value) (rerr error) {
	// We can pass anything to this function, which can cause the system panic. It's a
	// good idea to have a recovery from a panic.
	defer func() {
		if r := recover(); r != nil {
			rerr = fmt.Errorf("panic in append: %v", r)
		}
	}()

	// A few key conditional checks on the parameters. The above recover handler will handle edge cases.
	if sliceVal.Kind() != reflect.Ptr {
		return errors.New("values must be addressable so that we can set a value")
	}

	// What's the value the pointer is referencing.
	sliceVal = sliceVal.Elem()

	if sliceVal.Kind() != reflect.Slice {
		return errors.New("sliceVal must be a slice")
	}

	if itemVal.Kind() == reflect.Slice {
		sliceVal.Set(itemVal)
	} else {
		sliceVal.Set(reflect.Append(sliceVal, itemVal))
	}

	return nil
}

func main() {
	var names []string
	sliceVal := reflect.ValueOf(&names)

	if err := setToSlice(sliceVal, reflect.ValueOf("Peter")); err != nil {
		fmt.Println(err)
	}
	fmt.Println(names)

	if err := setToSlice(sliceVal, reflect.ValueOf([]string{"Stephanie", "Nancy"})); err != nil {
		fmt.Println(err)
	}
	fmt.Println(names)

	class := Class{}
	studentsVal := reflect.ValueOf(&class).Elem().FieldByName("Students").Addr()

	if err := setToSlice(studentsVal, reflect.ValueOf("Peter")); err != nil {
		fmt.Println(err)
	}
	fmt.Println(class)
}
