package ium

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
)

func Set[T any](column string, arg any) sq.QueryMod[T] {
	return SetEC[T](expr.JS(" = ", expr.S(column), expr.Arg(arg)))
}

func SetAN[T any](column string, argumentName string) sq.QueryMod[T] {
	return SetEC[T](expr.JS(" = ", expr.S(column), expr.ArgNamed(argumentName)))
}

func SetE[T any](column string, value litsql.Expression) sq.QueryMod[T] {
	return SetEC[T](expr.JS(" = ", expr.S(column), value))
}

func SetQ[T, A any](column string, q isq.Query[A]) sq.QueryMod[T] {
	return SetE[T](column, q)
}

func SetS[T any](column string, right string) sq.QueryMod[T] {
	return SetE[T](column, expr.S(right))
}

func SetC[T any](query string, args ...any) sq.QueryMod[T] {
	return SetEC[T](expr.C(query, args...))
}

func SetEC[T any](assignment litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Set{
			Set:     []litsql.Expression{assignment},
			Starter: true,
		})
	})
}
