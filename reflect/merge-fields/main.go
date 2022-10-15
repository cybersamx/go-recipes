package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// To keep this program simple, only string and int are supported.

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age,omitempty"`
	ZipCode   string `json:"zip_code,omitempty"`
}

func merge(dict map[string]any, user *User) *User {
	clone := *user

	typ := reflect.TypeOf(*user)
	val := reflect.ValueOf(&clone)

	for k, v := range dict {
		for i := 0; i < typ.NumField(); i++ {
			fieldTyp := typ.Field(i)
			fieldVal := val.Elem().FieldByName(fieldTyp.Name)

			tagVal := fieldTyp.Tag.Get("json")
			if strings.Contains(tagVal, k) {
				switch fieldVal.Kind() {
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
				default:
					log.Panicf("struct field tagged as %s has a unsupported type %v", k, fieldTyp)
				}
			}
		}
	}

	return &clone
}

func main() {
	userMap := map[string]any{
		"first_name": "Noel",
		"age":        30,
	}

	user := User{
		FirstName: "Gabriel",
		LastName:  "Lee",
		Age:       25,
		ZipCode:   "90405",
	}

	fmt.Println(user)
	merged := merge(userMap, &user)
	fmt.Println(merged)
}
