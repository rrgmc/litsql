package iim

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Into[T any](name string, columns ...string) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Table{
			Expression: expr.S(name),
			Columns:    columns,
		})
	})
}

func IntoAs[T any](name string, alias string, columns ...string) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Table{
			Expression: expr.S(name),
			Columns:    columns,
			Alias:      alias,
		})
	})
}
