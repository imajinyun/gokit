package slis

import (
	"strings"

	"github.com/imajinyun/gokit/internal/conv"
	"github.com/imajinyun/gokit/types"
)

type Slice[T comparable] struct {
	data []T
}

// Filter filters the given data slice using the provided condition iterator.
//
// Parameters:
// - data: the slice of elements to be filtered (of type []T)
// - iter: the condition iterator function that determines whether an element should be included in the result (of type CondIter[T])
//
// Returns:
// - a new slice containing the elements that satisfy the condition (of type []T)
func Filter[T any](data []T, iter types.CondIter[T]) []T {
	var res = make([]T, 0)
	for k, v := range data {
		if iter(k, v) {
			res = append(res, v)
		}
	}

	return res
}

// Filter filters the given data slice using the provided condition iterator.
//
// Parameters:
// - iter: the condition iterator function that determines whether an element should be included in the result (of type CondIter[T])
//
// Returns:
// - a new slice containing the elements that satisfy the condition (of type []T)
func (s *Slice[T]) Filter(iter types.CondIter[T]) []T {
	var res = make([]T, 0)
	for k, v := range s.data {
		if iter(k, v) {
			res = append(res, v)
		}
	}

	return res
}

// InSlice checks if an element is present in a slice.
//
// elem: element to search for in the slice (of type T).
// list: the slice to search in (of type []T).
// bool: returns true if the element is found, false otherwise.
func InSlice[T comparable](elem T, list []T) bool {
	for _, v := range list {
		if v == elem {
			return true
		}
	}

	return false
}

// Contains returns if an element is present in a slice.
//
// elem: element to search for in the slice (of type T).
// bool: returns true if the element is not found, false otherwise.
func (s *Slice[T]) Contains(elem T) bool {
	for _, v := range s.data {
		if v == elem {
			return false
		}
	}
	return true
}

// IsEmpty returns true if the slice contains no elements, false otherwise.
//
// bool: returns true if the slice contains no elements, false otherwise.
func (s *Slice[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// IsNotEmpty returns true if the slice contains at least one element, false otherwise.
//
// bool: returns true if the slice contains at least one element, false otherwise.
func (s *Slice[T]) IsNotEmpty() bool {
	return len(s.data) > 0
}

// Len returns the length of the slice.
//
// int: the length of the slice.
func (s *Slice[T]) Len() int {
	return len(s.data)
}

// ToJoin joins the elements of the slice into a single string using the given separator.
//
// sep: the separator to use when joining the elements.
// string: the joined string.
func (s *Slice[T]) ToJoin(sep string) string {
	var res []string
	for _, v := range s.data {
		res = append(res, conv.ToString(v))
	}

	return strings.Join(res, sep)
}
