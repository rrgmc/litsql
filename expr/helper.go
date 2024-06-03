package expr

import "github.com/rrgmc/litsql"

// ExprIf returns "e" if condition == true, nil otherwise.
func ExprIf(condition bool, e litsql.Expression) litsql.Expression {
	if condition {
		return e
	}
	return nil
}

// ExprIfElse returns "etrue" if condition == true, "efalse" otherwise.
func ExprIfElse(condition bool, etrue, efalse litsql.Expression) litsql.Expression {
	if condition {
		return etrue
	}
	return efalse
}

// PrefixIf returns an expression with the passed prefix if condition == true, only the expression if false.
func PrefixIf(condition bool, prefix litsql.Expression, e litsql.Expression) litsql.Expression {
	if condition {
		return J(prefix, e)
	}
	return e
}
