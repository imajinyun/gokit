package gohelper

import (
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// RandStr generates a random string of length n.
//
// Parameters:
// - n: the length of the random string.
//
// Return type:
// - string: the generated random string.
func RandStr(n int) string {
  if n > 1024 {
    n = 1024
  }

  var src = rand.NewSource(time.Now().UnixNano())
  r := rand.New(src)
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		b := byte(65 + r.Intn(25))
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

// Trim trims the specified characters from both ends of a string.
//
// Parameters:
// - s: the string to be trimmed.
// - char: the characters to be trimmed.
// Return type:
// - string: the trimmed string.
func Trim(s string, char string) string {
	return TrimLeft(TrimRight(s, char), char)
}

// TrimLeft removes leading characters from a string.
//
// Parameters:
// - s: the string to trim.
// - char: the characters to remove from the start of the string.
// Return type:
// - string: the trimmed string.
func TrimLeft(s string, char string) string {
	if char == "" {
		return strings.TrimLeftFunc(s, unicode.IsSpace)
	}

	r, _ := regexp.Compile("^[" + char + "]+")

	return r.ReplaceAllString(s, "")
}

// TrimRight removes trailing characters from a string.
//
// Parameters:
// - s: the string to be trimmed.
// - char: the characters to be trimmed.
// Return type:
// - string: the trimmed string.
func TrimRight(s string, char string) string {
	if char == "" {
		return strings.TrimRightFunc(s, unicode.IsSpace)
	}

	r, _ := regexp.Compile("[" + char + "]+$")

	return r.ReplaceAllString(s, "")
}
