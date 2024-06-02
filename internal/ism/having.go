package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Having[T any](condition string) sq.QueryMod[T] {
	return HavingE[T](expr.S(condition))
}

func HavingE[T any](condition litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Having{Conditions: []litsql.Expression{condition}})
	})
}

func HavingC[T any](query string, args ...any) sq.QueryMod[T] {
	return HavingE[T](expr.C(query, args...))
}
