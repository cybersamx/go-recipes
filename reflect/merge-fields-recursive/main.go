package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type City struct {
	Name    string `json:"name"`
	ZipCode string `json:"zipcode"`
}

type Location struct {
	Lat  float32 `json:"lat"`
	Long float32 `json:"long"`
	City City    `json:"city,omitempty"`
}

type User struct {
	Loc       Location `json:"loc,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Age       int      `json:"age,omitempty"`
}

// We can probably combine mergeReflect and merge functions together. Maybe refactor this
// later. Keep those functions separate to make it more readable.

func mergeReflect(dict map[string]any, typ reflect.Type, val reflect.Value) {
	for k, v := range dict {
		for i := 0; i < typ.NumField(); i++ {
			fieldTyp := typ.Field(i)
			fieldVal := val.FieldByName(fieldTyp.Name)

			tagVal := fieldTyp.Tag.Get("json")
			if strings.Contains(tagVal, k) {
				switch fieldTyp.Type.Kind() {
				case reflect.Struct:
					nestedDict, ok := v.(map[string]any)
					if !ok {
						log.Panicf("map[%s] with value %v must be of type map[string]any", k, v)
					}
					mergeReflect(nestedDict, fieldTyp.Type, fieldVal)
				case reflect.String:
					mapVal, ok := v.(string)
					if !ok {
						log.Panicf("map[%s] with value %v must be of type string", k, v)
					}
					fieldVal.SetString(mapVal)
				case reflect.Int:
					mapVal, ok := v.(int)
					if !ok {
						log.Panicf("map[%s] with value %v must be of type int", k, v)
					}
					fieldVal.SetInt(int64(mapVal))
				case reflect.Float32, reflect.Float64:
					mapVal, ok := v.(float64)
					if !ok {
						log.Panicf("map[%s] with value %v is not of type float64", k, v)
					}
					fieldVal.SetFloat(mapVal)
				default:
					log.Panicf("struct field tagged as %s has a unsupported type %v", k, fieldTyp)
				}
			}
		}
	}
}

func merge(dict map[string]any, obj any) {
	var typ reflect.Type
	if sf, ok := obj.(reflect.StructField); ok {
		typ = sf.Type
	} else {
		typ = reflect.TypeOf(obj)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
	}

	var val reflect.Value
	if rval, ok := obj.(reflect.Value); ok {
		val = rval
	} else {
		val = reflect.ValueOf(obj)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
	}

	mergeReflect(dict, typ, val)
}

func main() {
	userMap := map[string]any{
		"loc": map[string]any{
			"lat":  34.943,
			"long": -118.41,
			"city": map[string]any{
				"name":    "Los Angeles",
				"zipcode": "90012",
			},
		},
		"first_name": "Noel",
		"age":        30,
	}

	user := User{
		FirstName: "Gabriel",
		LastName:  "Lee",
		Age:       25,
		Loc: Location{
			Lat:  -73.935,
			Long: 40.731,
			City: City{
				Name:    "New York",
				ZipCode: "10001",
			},
		},
	}

	fmt.Println(user)
	merge(userMap, &user)
	fmt.Println(user)
}
