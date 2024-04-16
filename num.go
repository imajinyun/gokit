package gohelper

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

// RandInt generates a random integer within the specified range.
//
// min: the minimum value of the range
// max: the maximum value of the range
// int: the random integer within the specified range
func RandInt(min int, max int) int {
	if max-min <= 0 {
		panic("invalid range: max must be greater than min")
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return min + r.Intn(max-min)
}

// IsDigit checks if a string contains only digits.
//
// s: the string to be checked
// bool: returns true if the string contains only digits, false otherwise.
func IsDigit(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}

func IsNumber(s string) bool {
	re := regexp.MustCompile(`^-?\d+(\.\d+)?$`)

	return re.MatchString(s)
}

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)

	return err == nil
}

// IsInteger checks if the given value is an integer.
//
// T is the type of the value. It must be one of the following: int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64.
// value is the value to check.
// bool is the return type indicating if the value is an integer.
func IsInteger[T any](value T) bool {
	var v any = value

	switch v.(type) {
	case int, int8, int16, int32, int64:
		return true
	case uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64:
		return math.Remainder(reflect.ValueOf(v).Float(), 1) == 0
  case bool:
    return true
	case string:
		f, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			return false
		}
		return math.Remainder(f, 1) == 0
	default:
		panic(fmt.Sprintf("unhandled type: %T", value))
	}
}

func IsNatural[T any](value T) bool {
	return IsInteger(value) && IsPositive(value)
}

// IsPositive checks if the given value is positive.
//
// T is the type of the value.
// It must be one of the following: int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64.
// value is the value to check.
// bool is the return type indicating if the value is positive.
func IsPositive[T any](value T) (positive bool) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			panic(fmt.Sprintf("IsPositive panicked: %v", err))
		}
	}()

	var v any = value

	switch v.(type) {
	case int, int8, int16, int32, int64:
		positive = v.(int) >= 0
	case uint, uint8, uint16, uint32, uint64:
		positive = true
	case float32, float64:
		positive = math.Remainder(reflect.ValueOf(v).Float(), 1) == 0
  case bool:
    positive = If(v.(bool), true, false)
	case string:
		f, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			return false
		}
		_, frac := math.Modf(f)
		positive = frac == 0
	default:
		panic(fmt.Sprintf("unhandled type: %T", value))
	}

	return
}

// IsNegative checks if the given value is negative.
//
// T is the type of the value.
// It must be one of the following: int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64.
// value is the value to check.
// bool is the return type indicating if the value is negative.
func IsNegative[T any](value T) (negative bool) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			panic(fmt.Sprintf("IsNegative panicked: %v", err))
		}
	}()

	var v any = value

	switch v.(type) {
	case int, int8, int16, int32, int64:
		negative = v.(int) >= 0
	case uint, uint8, uint16, uint32, uint64:
		negative = true
	case float32, float64:
		negative = math.Remainder(reflect.ValueOf(v).Float(), 1) == 0
	case string:
		f, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			return false
		}
		_, frac := math.Modf(f)
		negative = frac == 0
	default:
		panic(fmt.Sprintf("unhandled type: %T", value))
	}

	return
}
