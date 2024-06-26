package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Distinct[T any](on ...string) sq.QueryMod[T] {
	return DistinctExpr[T](expr.StringList(on)...)
}

func DistinctExpr[T any](on ...litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Distinct{On: on})
	})
}
