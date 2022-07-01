package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Elevation struct {
	height float64 `sam:"height"`
	unit   string  `sam:"unit"`
}

type Coordinate struct {
	latitude  float64
	longitude float64
	elevation Elevation
}

type Location struct {
	Address    string
	City       string
	State      string
	Zip        string
	Coordinate Coordinate
}

// objToValues Return the values of obj and return the field names as a flatten string slice-basics.
func objToNames(obj interface{}, names []string, tag string) []string {
	var typ reflect.Type

	if sf, ok := obj.(reflect.StructField); ok {
		typ = sf.Type
	} else {
		typ = reflect.TypeOf(obj)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		if field.Type.Kind() == reflect.Struct {
			names = objToNames(field, names, tag)
			continue
		}

		// If tag is passed to the function, we only append if the field is tagged and that it matches tag.
		if tag == "" || field.Tag.Get(tag) != "" {
			names = append(names, field.Name)
		}
	}

	return names
}

// objToValues Return the values of obj and return the field values as a flatten string slice-basics.
func objToValues(obj interface{}, vals []string) []string {
	var val reflect.Value

	if rval, ok := obj.(reflect.Value); ok {
		val = rval
	} else {
		val = reflect.ValueOf(obj)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			vals = objToValues(field, vals)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			vals = append(vals, strconv.FormatInt(field.Int(), 10))
		case reflect.Float32, reflect.Float64:
			vals = append(vals, strconv.FormatFloat(field.Float(), 'f', 3, 32))
		case reflect.String:
			vals = append(vals, field.String())
		default:
			// This is sample code, so this is fine.
			// As a best practice, we should really have an exhaustive list of cases to extract the field value.
			vals = append(vals, "<Field type not supported>")
		}
	}

	return vals
}

func main() {
	loc := Location{
		Address: "1111 S. Figueroa St",
		City:    "Los Angeles",
		State:   "CA",
		Zip:     "90015",
		Coordinate: Coordinate{
			latitude:  34.043450,
			longitude: -118.266340,
			elevation: Elevation{
				height: 10,
				unit:   "ft",
			},
		},
	}

	names := objToNames(loc, make([]string, 0), "")
	fmt.Println("Flatten names")
	fmt.Println(names)

	taggedNames := objToNames(loc, make([]string, 0), "sam")
	fmt.Println("Flatten names that are tagged")
	fmt.Println(taggedNames)

	values := objToValues(loc, make([]string, 0))
	fmt.Println("Flatten values")
	fmt.Println(values)
}
