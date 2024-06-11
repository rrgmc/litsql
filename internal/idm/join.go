package idm

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
)

//litsql:dialects none
func JoinExpr[T, CHAIN any](typ string, table litsql.Expression) *ichain.JoinChain[T, CHAIN] {
	return ichain.NewJoinChain[T, CHAIN](&ichain.JoinChain[T, CHAIN]{
		Join: &iclause.Join{
			Type: typ,
			To:   &iclause.From{Table: table},
		},
	})
}

func InnerJoin[T, CHAIN any](table string) *ichain.JoinChain[T, CHAIN] {
	return InnerJoinExpr[T, CHAIN](expr.String(table))
}

func InnerJoinExpr[T, CHAIN any](table litsql.Expression) *ichain.JoinChain[T, CHAIN] {
	return JoinExpr[T, CHAIN](ichain.JoinInnerJoin, table)
}

func LeftJoin[T, CHAIN any](table string) *ichain.JoinChain[T, CHAIN] {
	return LeftJoinExpr[T, CHAIN](expr.String(table))
}

func LeftJoinExpr[T, CHAIN any](table litsql.Expression) *ichain.JoinChain[T, CHAIN] {
	return JoinExpr[T, CHAIN](ichain.JoinLeftJoin, table)
}

func RightJoin[T, CHAIN any](table string) *ichain.JoinChain[T, CHAIN] {
	return RightJoinExpr[T, CHAIN](expr.String(table))
}

func RightJoinExpr[T, CHAIN any](table litsql.Expression) *ichain.JoinChain[T, CHAIN] {
	return JoinExpr[T, CHAIN](ichain.JoinRightJoin, table)
}

func FullJoin[T, CHAIN any](table string) *ichain.JoinChain[T, CHAIN] {
	return FullJoinExpr[T, CHAIN](expr.String(table))
}

func FullJoinExpr[T, CHAIN any](table litsql.Expression) *ichain.JoinChain[T, CHAIN] {
	return JoinExpr[T, CHAIN](ichain.JoinFullJoin, table)
}

func CrossJoin[T, CHAIN any](table string) *ichain.JoinChain[T, CHAIN] {
	return CrossJoinExpr[T, CHAIN](expr.String(table))
}

func CrossJoinExpr[T, CHAIN any](table litsql.Expression) *ichain.JoinChain[T, CHAIN] {
	return JoinExpr[T, CHAIN](ichain.JoinCrossJoin, table)
}

func StraightJoin[T, CHAIN any](table string) *ichain.JoinChain[T, CHAIN] {
	return StraightJoinExpr[T, CHAIN](expr.String(table))
}

func StraightJoinExpr[T, CHAIN any](table litsql.Expression) *ichain.JoinChain[T, CHAIN] {
	return JoinExpr[T, CHAIN](ichain.JoinStraightJoin, table)
}
