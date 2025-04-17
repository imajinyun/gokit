package types

import (
	"github.com/imajinyun/gokit/internal/datetime"
)

type DateTime struct {
	*datetime.DateTime
}



type CondIter[T any] func(k int, v T) bool

type Option struct {
	IncludeNumber    bool
	IncludeUppercase bool
	IncludeSymbol    bool
}
