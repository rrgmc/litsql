package idm

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Where[T any](condition string) sq.QueryMod[T] {
	return WhereExpr[T](expr.String(condition))
}

func WhereExpr[T any](condition litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Where{Conditions: []litsql.Expression{condition}})
	})
}

func WhereClause[T any](query string, args ...any) sq.QueryMod[T] {
	return WhereExpr[T](expr.Clause(query, args...))
}
