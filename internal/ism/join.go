package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq/chain"
)

func InnerJoin[T any](table string) chain.Join[T] {
	return InnerJoinE[T](expr.String(table))
}

func InnerJoinE[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinInnerJoin, To: &iclause.From{Table: table}}}
}

func LeftJoin[T any](table string) chain.Join[T] {
	return LeftJoinE[T](expr.String(table))
}

func LeftJoinE[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinLeftJoin, To: &iclause.From{Table: table}}}
}

func RightJoin[T any](table string) chain.Join[T] {
	return RightJoinE[T](expr.String(table))
}

func RightJoinE[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinRightJoin, To: &iclause.From{Table: table}}}
}

func FullJoin[T any](table string) chain.Join[T] {
	return FullJoinE[T](expr.String(table))
}

func FullJoinE[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinFullJoin, To: &iclause.From{Table: table}}}
}

func CrossJoin[T any](table string) chain.Join[T] {
	return CrossJoinE[T](expr.String(table))
}

func CrossJoinE[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinCrossJoin, To: &iclause.From{Table: table}}}
}

func StraightJoin[T any](table string) chain.Join[T] {
	return StraightJoinE[T](expr.String(table))
}

func StraightJoinE[T any](table litsql.Expression) chain.Join[T] {
	return &ichain.JoinChain[T]{Join: &iclause.Join{Type: ichain.JoinStraightJoin, To: &iclause.From{Table: table}}}
}
