package marshaler

import (
	"encoding/json"
	"reflect"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// CREDIT: Go standard library encoding/json.
// https://github.com/golang/go/blob/64c9ee98b7684cf2156f620cbab4dbb6081b9771/src/encoding/json/encode.go#L339
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func toMap(user UserReflect) (map[string]interface{}, error) {
	typ := reflect.TypeOf(user)
	val := reflect.ValueOf(user)

	dict := make(map[string]interface{}, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == "Password" {
			continue
		}

		tag := field.Tag
		jsonTag := tag.Get("json")
		jsonVals := strings.Split(jsonTag, ",")
		name := jsonVals[0]
		if name == "-" {
			continue
		}
		if contains(jsonVals[1:], "omitempty") && isEmptyValue(val.Field(i)) {
			continue
		}

		dict[name] = val.Field(i).Interface()
	}

	return dict, nil
}

type UserReflect struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Note      string    `json:"note,omitempty"`
	Tags      string    `json:"-"`
}

func (u UserReflect) MarshalJSON() ([]byte, error) {
	userMap, err := toMap(u)
	if err != nil {
		return nil, err
	}

	return json.Marshal(userMap)
}
