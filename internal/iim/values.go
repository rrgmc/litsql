package iim

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
)

func Values[T any](values ...any) sq.QueryMod[T] {
	return ValuesExpr[T](expr.Args(values)...)
}

func ValuesArgNamed[T any](argumentNames ...string) sq.QueryMod[T] {
	return ValuesExpr[T](expr.ArgsNamed(argumentNames...)...)
}

func ValuesExpr[T any](clauses ...litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		var vals []iclause.Value
		if len(clauses) > 0 {
			vals = append(vals, clauses)
		}
		a.AddQueryClause(&iclause.Values{
			Vals: vals,
		})
	})
}

func ValuesString[T any](clauses ...string) sq.QueryMod[T] {
	return ValuesExpr[T](expr.StringList(clauses)...)
}

func Query[T, A any](q isq.Query[A]) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Values{
			Query: q,
		})
	})
}
