package strs

import (
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

const lowercs = "abcdefghijklmnopqrstuvwxyz"
const uppercs = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const symbols = `~!@#$%^&*()_+{}|:"<>?-=[]\;',./`

const letterIdxBits = 6
const letterIdxMask = 2<<letterIdxBits - 1
const letterIdxMax = 63 / letterIdxBits

type Option struct {
	IncludeNumber    bool
	IncludeUppercase bool
	IncludeSymbol    bool
}

// RandStr generates a random string of length n.
//
// n: the length of the string to be generated.
// string: the randomly generated string.
//
// @see https://github.com/tamboto2000/random
func RandStr(n int) string {
	return gen(n, Option{})
}

// RandStrWithOption generates a random string of length n with custom options.
//
// n: the length of the string to be generated.
// opt: the options to include in the generated string.
// string: the randomly generated string.
func RandStrWithOption(n int, opt Option) string {
	return gen(n, opt)
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

func gen(n int, opt Option) string {
	letters := lowercs

	if opt.IncludeNumber {
		letters += numbers
	}

	if opt.IncludeUppercase {
		letters += uppercs
	}

	if opt.IncludeSymbol {
		letters += symbols
	}

	src := rand.NewSource(time.Now().UnixNano())
	byt := make([]byte, n)

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letters) {
			byt[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&byt))
}
