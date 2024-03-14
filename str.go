package gohelper

import "math/rand"

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
