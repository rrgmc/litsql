package expr

import "github.com/rrgmc/litsql"

func ExprIf(condition bool, e litsql.Expression) litsql.Expression {
	if condition {
		return e
	}
	return nil
}

func ExprIfElse(condition bool, etrue, efalse litsql.Expression) litsql.Expression {
	if condition {
		return etrue
	}
	return efalse
}

func PrefixIf(condition bool, prefix litsql.Expression, e litsql.Expression) litsql.Expression {
	if condition {
		return J(prefix, e)
	}
	return e
}
