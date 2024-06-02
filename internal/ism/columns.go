package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Columns[T any](names ...string) sq.QueryMod[T] {
	return ColumnsE[T](expr.SL(names)...)
}

func ColumnsE[T any](names ...litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Columns{Columns: names})
	})
}
