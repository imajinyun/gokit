package gohelper

import (
	"crypto/rand"
	"math/big"
)

func RandNum(min int, max int) int {
	v := max - min
	if v <= 0 {
		panic("invalid range: max must be greater than min")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(v)))
	if err != nil {
		panic(err)
	}

	return int(n.Int64()) + min
}
