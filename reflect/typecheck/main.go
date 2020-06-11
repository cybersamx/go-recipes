package main

import (
	"fmt"
	"reflect"
)

func typeSwitch(vals []interface{}) {
	fmt.Println("Type check")
	for _, v := range vals {
		switch actual := v.(type) {
		case int:
			fmt.Println(actual, "is int")
		case float64:
			fmt.Println(actual, "is float64")
		case float32:
			fmt.Println(actual, "is float32")
		case *int:
			fmt.Println(actual, "is *int")
		case string:
			fmt.Println(actual, "is string")
		default:
			fmt.Println(actual, "is unsupported type")
		}
	}
	fmt.Println()
}

func typeCast(vals []interface{}) {
	// Recover from panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
			fmt.Println()
		}
	}()

	fmt.Println("Type cast")
	for _, v := range vals {
		if _, ok := v.(int); ok {
			fmt.Println(v, "is int")
		}
		if _, ok := v.(float64); ok {
			fmt.Println(v, "is float64")
		}
		if _, ok := v.(float32); ok {
			fmt.Println(v, "is float32")
		}
		if _, ok := v.(*int); ok {
			fmt.Println(v, "is *int")
		}
		if _, ok := v.(string); ok {
			fmt.Println(v, "is string")
		}
	}

	// When typecasting in Go, always include the second variable. If you cast
	// a value with only 1 variable, it's okay if the type assertion succeeds.
	var v interface{} = "hello world"
	str := v.(string)   // No problem.
	fmt.Println(str)

	// But if the type assertion fails, the program will panic.
	n := v.(int)    // Panic
	fmt.Println(n)
}

func reflectKind(vals []interface{}) {
	fmt.Println("Type cast")
	for _, v := range vals {
		typ := reflect.TypeOf(v)
		switch typ.Kind() {
		case reflect.Int:
			fmt.Println(v, "is int")
		case reflect.Float64:
			fmt.Println(v, "is float64")
		case reflect.Float32:
			fmt.Println(v, "is float32")
		case reflect.Ptr:
			fmt.Println(v, "is pointer")
		case reflect.String:
			fmt.Println(v, "is string")
		default:
			fmt.Println(v, "is unsupported type")
		}
	}
	fmt.Println()
}

func main() {
	var n int = 1

	vals := []interface{}{
		2.3,
		float32(4.5),
		1,
		"text",
		&n,
	}

	typeSwitch(vals)
	typeCast(vals)
	reflectKind(vals)
}
