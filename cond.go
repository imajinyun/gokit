package gokit

import "github.com/imajinyun/gokit/internal/cond"

// If is a function that takes a boolean expression and two values, and returns one of the values based on the evaluation of the expression.
//
// expr: a boolean expression.
// succ: the value to return if the expression is true.
// fail: the value to return if the expression is false.
// T: the type of the returned value.
//
// Returns:
// The value of succ if expr is true, otherwise the value of fail.
func If[T any](expr bool, succ T, fail T) T {
	return cond.If(expr, succ, fail)
}
