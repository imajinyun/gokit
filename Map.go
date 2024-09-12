package gohelper

// GetOrDefault retrieves the value from a map based on the key. If the key does not exist, it returns the default value.
//
// Parameters:
// m - the map to retrieve the value from (of type map[K]V)
// key - the key to look up in the map (of type K)
// def - the default value to return if the key is not found (of type V)
// Return type: V
func GetOrDefault[K comparable, V any](m map[K]V, key K, def V) V {
	if val, ok := m[key]; ok {
		return val
	}

	return def
}
