package ism

import (
	"fmt"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Limit[T any](count int) sq.QueryMod[T] {
	return LimitE[T](expr.S(fmt.Sprint(count)))
}

func LimitE[T any](count litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Limit{Count: count})
	})
}

func LimitA[T any](arg any) sq.QueryMod[T] {
	return LimitE[T](expr.C("?", arg))
}
