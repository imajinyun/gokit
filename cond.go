package gohelper

func If(expr bool, succ any, fail any) any {
	if expr {
		return succ
	}

	return fail
}
