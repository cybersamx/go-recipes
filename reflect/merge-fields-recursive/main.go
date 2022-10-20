package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

type City struct {
	Name    string `json:"name"`
	ZipCode string `json:"zipcode"`
}

type Location struct {
	Lat  float32 `json:"lat"`
	Long float64 `json:"long"`
	City City    `json:"city,omitempty"`
}

type User struct {
	Loc       Location `json:"loc,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Age       int      `json:"age,omitempty"`
	NumKids   int16    `json:"num_kids,omitempty"`
	UpdatedAt int64    `json:"updated_at,omitempty"`
}

// Credit: https://github.com/golang/go/blob/master/src/html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved. Licensed under BSD.
func indirect(v any) any {
	if v == nil {
		return nil
	}
	if rtyp := reflect.TypeOf(v); rtyp.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return v
	}

	rval := reflect.ValueOf(v)
	for rval.Kind() == reflect.Ptr && !rval.IsNil() {
		rval = rval.Elem()
	}

	return rval.Interface()
}

func toInt64(v any) (int64, error) {
	val := indirect(v)

	switch i := val.(type) {
	case int64:
		return i, nil
	case int:
		return int64(i), nil
	case int32:
		return int64(i), nil
	case int16:
		return int64(i), nil
	case int8:
		return int64(i), nil
	case float64:
		return int64(math.Round(i)), nil
	case float32:
		return int64(math.Round(float64(i))), nil
	default:
		return 0, fmt.Errorf("can't cast %v of type %T to int64", val, val)
	}
}

func toFloat64(v any) (float64, error) {
	val := indirect(v)
	
	switch i := val.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int16:
		return float64(i), nil
	case int8:
		return float64(i), nil
	default:
		return 0.0, fmt.Errorf("can't cast %v of type %T to int64", val, val)
	}
}

func mergeReflect(dict map[string]any, rval reflect.Value) error {
	rtyp := rval.Type()

	for k, v := range dict {
		for i := 0; i < rtyp.NumField(); i++ {
			fieldTyp := rtyp.Field(i)
			fieldVal := rval.FieldByName(fieldTyp.Name)

			tagVal := fieldTyp.Tag.Get("json")
			if strings.Contains(tagVal, k) {
				switch fieldTyp.Type.Kind() {
				case reflect.Struct:
					mapVal, ok := v.(map[string]any)
					if !ok {
						return fmt.Errorf("map[%s] with value %v must be of type map[string]any", k, v)
					}
					if err := mergeReflect(mapVal, fieldVal); err != nil {
						return err
					}
				case reflect.String:
					stringVal, ok := v.(string)
					if !ok {
						return fmt.Errorf("map[%s] with value %v must be of type string", k, v)
					}
					fieldVal.SetString(stringVal)
				case reflect.Float64, reflect.Float32:
					val, err := toFloat64(v)
					if err != nil {
						return fmt.Errorf("map[%s] with value %v is not of type float*", k, v)
					}
					fieldVal.SetFloat(val) // SetInt will take care of casting to the right precision ie. Float32, Float64
				case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
					val, err := toInt64(v)
					if err != nil {
						return fmt.Errorf("map[%s] with value %v is not of type int*", k, v)
					}
					fieldVal.SetInt(val) // SetInt will take care of casting to the right precision ie. Int8, Int32, etc.
				default:
					return fmt.Errorf("struct field tagged as %s has a unsupported type %v", k, fieldTyp)
				}
			}
		}
	}

	return nil
}

func merge(dict map[string]any, obj any) error {
	rval := reflect.ValueOf(obj)
	for rval.Kind() == reflect.Ptr && !rval.IsNil() {
		rval = rval.Elem()
	}

	return mergeReflect(dict, rval)
}

func main() {
	userMap := map[string]any{
		"loc": map[string]any{
			"lat":  float32(34.943),
			"long": -118.41,
			"city": map[string]any{
				"name":    "Los Angeles",
				"zipcode": "90012",
			},
		},
		"first_name": "Noel",
		"age":        30,
		"num_kids":   int16(3),
		"updated_at": 1666241681,
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
	if err := merge(userMap, &user); err != nil {
		panic(err)
	}
	fmt.Println(user)
}
