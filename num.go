package gohelper

import (
	"math/rand"
	"time"
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
