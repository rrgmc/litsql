package iim

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
)

func Values[T any](values ...any) sq.QueryMod[T] {
	return ValuesE[T](expr.Args(values)...)
}

func ValuesE[T any](clauses ...litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		var vals []iclause.Value
		if len(clauses) > 0 {
			vals = append(vals, clauses)
		}
		a.Add(&iclause.Values{
			Vals: vals,
		})
	})
}

func ValuesS[T any](clauses ...string) sq.QueryMod[T] {
	return ValuesE[T](expr.SL(clauses)...)
}

// Insert from a query
func Query[T, A any](q isq.Query[A]) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Values{
			Query: q,
		})
	})
}
