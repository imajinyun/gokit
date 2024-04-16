package gohelper

type CondIter[T any] func(k int, v T) bool

// Filter filters the given data slice using the provided condition iterator.
//
// Parameters:
// - data: the slice of elements to be filtered (of type []T)
// - iter: the condition iterator function that determines whether an element should be included in the result (of type CondIter[T])
//
// Returns:
// - a new slice containing the elements that satisfy the condition (of type []T)
func Filter[T any](data []T, iter CondIter[T]) []T {
	var res = make([]T, 0)
	for k, v := range data {
		if iter(k, v) {
			res = append(res, v)
		}
	}

	return res
}
