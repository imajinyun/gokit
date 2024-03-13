package gohelper

import (
	"reflect"
	"strings"
)

func StructToMap(data any) map[string]any {
	r := make(map[string]any)
	v := reflect.ValueOf(data)

	for i := 0; i < v.NumField(); i++ {
		k := v.Type().Field(i).Tag.Get("json")
		if k == "" {
			k = v.Type().Field(i).Name
		}
		v := v.Field(i).Interface()
		r[strings.ToLower(k)] = v
	}

	return r
}
