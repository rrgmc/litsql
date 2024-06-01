package expr

import "github.com/rrgmc/litsql"

func CastSlice[E litsql.Expression](list []E) []litsql.Expression {
	var ret []litsql.Expression
	for _, e := range list {
		ret = append(ret, e)
	}
	return ret
}
