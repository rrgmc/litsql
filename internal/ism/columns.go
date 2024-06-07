package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Columns[T any](names ...string) sq.QueryMod[T] {
	return ColumnsE[T](expr.StringList(names)...)
}

func ColumnsE[T any](names ...litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Columns{Columns: names})
	})
}

func ColumnsC[T any](query string, args ...any) sq.QueryMod[T] {
	return ColumnsE[T](expr.Clause(query, args...))
}
