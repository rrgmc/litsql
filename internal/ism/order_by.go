package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func OrderBy[T any](names ...string) sq.QueryMod[T] {
	return OrderByE[T](expr.SL(names)...)
}

func OrderByE[T any](names ...litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.OrderBy{Expressions: names})
	})
}
