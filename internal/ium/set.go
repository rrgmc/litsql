package ium

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
)

func Set[T any](column string, arg any) sq.QueryMod[T] {
	return SetExprClause[T](expr.JoinSep(" = ", expr.String(column), expr.Arg(arg)))
}

func SetArgNamed[T any](column string, argumentName string) sq.QueryMod[T] {
	return SetExprClause[T](expr.JoinSep(" = ", expr.String(column), expr.ArgNamed(argumentName)))
}

func SetExpr[T any](column string, value litsql.Expression) sq.QueryMod[T] {
	return SetExprClause[T](expr.JoinSep(" = ", expr.String(column), value))
}

func SetQuery[T, A any](column string, q isq.Query[A]) sq.QueryMod[T] {
	return SetExpr[T](column, q)
}

func SetString[T any](column string, right string) sq.QueryMod[T] {
	return SetExpr[T](column, expr.String(right))
}

func SetClause[T any](query string, args ...any) sq.QueryMod[T] {
	return SetExprClause[T](expr.Clause(query, args...))
}

func SetExprClause[T any](assignment litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Set{
			Set:     []litsql.Expression{assignment},
			Starter: true,
		})
	})
}
