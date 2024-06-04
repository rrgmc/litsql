package expr

import "github.com/rrgmc/litsql"

// CastSlice converts a slice of Expression-implemented items to a slice of [litsql.Expression].
func CastSlice[E litsql.Expression](list []E) []litsql.Expression {
	var ret []litsql.Expression
	for _, e := range list {
		ret = append(ret, e)
	}
	return ret
}

// MapSlice converts a slice of data to a slice of [litsql.Expression] using a map function.
func MapSlice[T any](source []T, mapFunc func(T) litsql.Expression) []litsql.Expression {
	ret := make([]litsql.Expression, len(source))
	for i, val := range source {
		ret[i] = mapFunc(val)
	}
	return ret
}
