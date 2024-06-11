package idm

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/isq"
)

func Using[T, CHAIN any](table string) *ichain.FromChain[T, CHAIN] {
	return UsingExpr[T, CHAIN](expr.String(table))
}

func UsingExpr[T, CHAIN any](table litsql.Expression) *ichain.FromChain[T, CHAIN] {
	return ichain.NewFromChain[T, CHAIN](&ichain.FromChain[T, CHAIN]{
		From: &iclause.From{
			Table:   table,
			Starter: true,
			Clause:  "USING",
		},
	})
}

func UsingQuery[T, CHAIN, A any](q isq.Query[A]) *ichain.FromChain[T, CHAIN] {
	return UsingExpr[T, CHAIN](q)
}
