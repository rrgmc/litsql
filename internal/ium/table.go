package ium

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Table[T any](name string) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Table{
			Expression: expr.S(name),
		})
	})
}

func TableAs[T any](name string, alias string) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Table{
			Expression: expr.S(name),
			Alias:      alias,
		})
	})
}
