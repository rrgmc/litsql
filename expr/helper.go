package expr

import "github.com/rrgmc/litsql"

// If returns "e" if condition == true, nil otherwise.
func If(condition bool, e litsql.Expression) litsql.Expression {
	return IfElse(condition, e, nil)
}

// IfElse returns "etrue" if condition == true, "efalse" otherwise.
func IfElse(condition bool, etrue, efalse litsql.Expression) litsql.Expression {
	if condition {
		return etrue
	}
	return efalse
}

// PrefixIf returns an expression with the passed prefix if condition == true, only the expression if false.
func PrefixIf(condition bool, prefix litsql.Expression, e litsql.Expression) litsql.Expression {
	if condition {
		return Join(prefix, e)
	}
	return e
}
