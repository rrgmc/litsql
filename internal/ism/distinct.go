package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Distinct[T any](on ...string) sq.QueryMod[T] {
	return DistinctE[T](expr.SL(on)...)
}

func DistinctE[T any](on ...litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Distinct{On: on})
	})
}
