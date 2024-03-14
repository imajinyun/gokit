package gohelper

// If is a function that takes a boolean expression and two values, and returns one of the values based on the evaluation of the expression.
//
// expr: a boolean expression.
// succ: the value to return if the expression is true.
// fail: the value to return if the expression is false.
// any: the type of the returned value.
func If(expr bool, succ any, fail any) any {
	if expr {
		return succ
	}

	return fail
}
