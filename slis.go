package gokit

import (
	"github.com/imajinyun/gohelper/types"
)

func Filter[T any](data []T, iter types.CondIter[T]) []T {
	var res = make([]T, 0)
	for k, v := range data {
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
