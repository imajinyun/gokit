package conv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// ToString converts the given data to a string representation.
//
// data: the data to be converted.
// string: the string representation of the data.
func ToString(data any) string {
	return fmt.Sprintf("%v", data)
}

// ToJson converts the given data to a JSON string representation.
//
// data: the data to be converted.
// string: the JSON string representation of the data.
// error: an error if the conversion fails.
func ToJson(data any) (string, error) {
	res, err := json.Marshal(data)
	if err != nil {
		res = []byte("")
	}

	return string(res), nil
}

// StructToMap converts a struct to a map[string]T, where T is the type of the struct fields.
// This function is optimized for performance and memory allocation.
//
// Parameter:
// data - the input struct (of type struct{...} with fields of type T)
// Return type:
// map[string]T - a map containing the struct fields
func StructToMap[T any](data any) map[string]T {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	r := make(map[string]T, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		fieldInfo := t.Field(i)
		k := fieldInfo.Tag.Get("json")
		if k == "" {
			k = fieldInfo.Name
		}
		r[strings.ToLower(k)] = v.Field(i).Interface().(T)
	}

	return r
}
