package gokit

import (
	"github.com/imajinyun/gokit/internal/slis"
	"github.com/imajinyun/gokit/types"
)

func Filter[T any](data []T, iter types.CondIter[T]) []T {
	return slis.Filter(data, iter)
}

func InSlice[T comparable](elem T, list []T) bool {
	return slis.InSlice(elem, list)
}
