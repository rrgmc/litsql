package ium

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq/chain"
)

func InnerJoin[T any](table string) chain.Join[T] {
	return InnerJoinExpr[T](expr.String(table))
}

func InnerJoinExpr[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinInnerJoin, To: &iclause.From{Table: table}}}
}

func LeftJoin[T any](table string) chain.Join[T] {
	return LeftJoinExpr[T](expr.String(table))
}

func LeftJoinExpr[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinLeftJoin, To: &iclause.From{Table: table}}}
}

func RightJoin[T any](table string) chain.Join[T] {
	return RightJoinExpr[T](expr.String(table))
}

func RightJoinExpr[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinRightJoin, To: &iclause.From{Table: table}}}
}

func FullJoin[T any](table string) chain.Join[T] {
	return FullJoinExpr[T](expr.String(table))
}

func FullJoinExpr[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinFullJoin, To: &iclause.From{Table: table}}}
}

func CrossJoin[T any](table string) chain.Join[T] {
	return CrossJoinExpr[T](expr.String(table))
}

func CrossJoinExpr[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinCrossJoin, To: &iclause.From{Table: table}}}
}

func StraightJoin[T any](table string) chain.Join[T] {
	return StraightJoinExpr[T](expr.String(table))
}

func StraightJoinExpr[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinStraightJoin, To: &iclause.From{Table: table}}}
}
