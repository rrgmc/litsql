package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

func Offset[T any](count int) sq.QueryMod[T] {
	return OffsetExpr[T](expr.Arg(count))
}

func OffsetExpr[T any](count litsql.Expression) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.Offset{Count: count})
	})
}

func OffsetArg[T any](arg any) sq.QueryMod[T] {
	return OffsetExpr[T](expr.Arg(arg))
}

func OffsetArgNamed[T any](argumentName string) sq.QueryMod[T] {
	return OffsetExpr[T](expr.ArgNamed(argumentName))
}
