package gohelper

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
