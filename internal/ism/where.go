package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Where[T any](condition string) sq.QueryMod[T] {
	return WhereE[T](expr.S(condition))
}

func WhereE[T any](condition litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Where{Conditions: []litsql.Expression{condition}})
	})
}

func WhereC[T any](query string, args ...any) sq.QueryMod[T] {
	return WhereE[T](expr.C(query, args...))
}
