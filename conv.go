package gohelper

import (
	"reflect"
	"strings"
)

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
