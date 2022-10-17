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
	Enable     bool
	Coordinate Coordinate
}

// objToNames Return the values of obj and return the field names as a flatten string slice-custom-type.
func objToNames(obj any, names []string, tag string) []string {
	rtyp := reflect.TypeOf(obj)
	if rtyp.Kind() == reflect.Ptr {
		rtyp = rtyp.Elem()
	}

	return rtypToNames(rtyp, names, tag)
}

func rtypToNames(rtyp reflect.Type, names []string, tag string) []string {
	for i := 0; i < rtyp.NumField(); i++ {
		field := rtyp.Field(i)

		if field.Type.Kind() == reflect.Struct {
			names = rtypToNames(field.Type, names, tag)
			continue
		}

		// Append the field name if no tag is passed to the function. But if a tag is
		// passed, we only append if there's a non-empty value found for that tag.
		if tag == "" || field.Tag.Get(tag) != "" {
			names = append(names, field.Name)
		}
	}

	return names
}

// objToValues Return the values of obj and return the field values as a flatten string slice-custom-type.
func objToValues(obj any, vals []string) []string {
	rval := reflect.ValueOf(obj)
	if rval.Kind() == reflect.Ptr {
		rval = rval.Elem()
	}

	return rvalToValues(rval, vals)
}

func rvalToValues(rval reflect.Value, vals []string) []string {
	for i := 0; i < rval.NumField(); i++ {
		field := rval.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			vals = rvalToValues(field, vals)
		case reflect.Int:
			vals = append(vals, strconv.FormatInt(field.Int(), 10))
		case reflect.Float32, reflect.Float64:
			vals = append(vals, strconv.FormatFloat(field.Float(), 'f', 3, 32))
		case reflect.String:
			vals = append(vals, field.String())
		default:
			// This is sample code, so this is fine.
			// As a best practice, we should really have an exhaustive list of cases to extract the field value.
			vals = append(vals, fmt.Sprintf("<unsupported-type: %v>", field.Kind()))
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

	names := objToNames(loc, []string{}, "")
	fmt.Println("Flatten names")
	fmt.Println(names)

	taggedNames := objToNames(loc, []string{}, "sam")
	fmt.Println("Flatten names that are tagged")
	fmt.Println(taggedNames)

	values := objToValues(loc, []string{})
	fmt.Println("Flatten values")
	fmt.Println(values)
}
