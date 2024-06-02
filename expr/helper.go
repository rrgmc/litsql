package expr

import "github.com/rrgmc/litsql"

func ExprIf(condition bool, e litsql.Expression) litsql.Expression {
	if condition {
		return e
	}
	return nil
}

func PrefixIf(condition bool, prefix litsql.Expression, e litsql.Expression) litsql.Expression {
	if condition {
		return J(prefix, e)
	}
	return e
}
