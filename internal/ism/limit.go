package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Limit[T any](count int) sq.QueryMod[T] {
	return LimitE[T](expr.Arg(count))
}

func LimitE[T any](count litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Limit{Count: count})
	})
}

func LimitA[T any](arg any) sq.QueryMod[T] {
	return LimitE[T](expr.Arg(arg))
}

func LimitAN[T any](argumentName string) sq.QueryMod[T] {
	return LimitE[T](expr.ArgNamed(argumentName))
}
