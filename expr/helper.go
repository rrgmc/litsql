package expr

import "github.com/rrgmc/litsql"

func PrefixIf(condition bool, prefix litsql.Expression, e litsql.Expression) litsql.Expression {
	if condition {
		return J(prefix, e)
	}
	return e
}
