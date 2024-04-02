package gohelper

import (
	"math/rand"
	"reflect"
)

// RandStr generates a random string of a specified length.
//
// len int - the length of the random string.
// Returns string - the randomly generated string.
func RandStr(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := byte(65 + rand.Intn(25))
		bytes[i] = b
	}

	return string(bytes)
}

func Empty(val any) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Slice:
		return v.Len() == 0 || (v.Cap() != 0 && v.IsNil())
	case reflect.Map:
		return v.Len() == 0 || v.IsNil()
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

	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}
