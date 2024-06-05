package ism

import (
	"fmt"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Offset[T any](count int) sq.QueryMod[T] {
	return OffsetE[T](expr.S(fmt.Sprint(count)))
}

func OffsetE[T any](count litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Offset{Count: count})
	})
}

func OffsetA[T any](arg any) sq.QueryMod[T] {
	return OffsetE[T](expr.Arg(arg))
}

func OffsetAN[T any](argumentName string) sq.QueryMod[T] {
	return OffsetE[T](expr.ArgNamed(argumentName))
}
