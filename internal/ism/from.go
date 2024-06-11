package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/isq"
)

func From[T, CHAIN any](table string) *ichain.FromChain[T, CHAIN] {
	return FromExpr[T, CHAIN](expr.String(table))
}

func FromExpr[T, CHAIN any](table litsql.Expression) *ichain.FromChain[T, CHAIN] {
	return ichain.NewFromChain[T, CHAIN](&ichain.FromChain[T, CHAIN]{
		From: &iclause.From{
			Table:   table,
			Starter: true,
		},
	})
}

func FromQuery[T, CHAIN, A any](q isq.Query[A]) *ichain.FromChain[T, CHAIN] {
	return FromExpr[T, CHAIN](q)
}
