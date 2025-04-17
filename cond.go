package gokit

import "github.com/imajinyun/gokit/internal/cond"

func If[T any](expr bool, succ T, fail T) T {
	return cond.If(expr, succ, fail)
}
