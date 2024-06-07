package idm

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq/chain"
)

func Using[T any](table string) chain.From[T] {
	return UsingE[T](expr.String(table))
}

func UsingE[T any](table litsql.Expression) chain.From[T] {
	return &ichain.FromChain[T]{From: &iclause.From{Table: table, Starter: true, Clause: "USING"}}
}

func UsingQ[T, A any](q isq.Query[A]) chain.From[T] {
	return UsingE[T](q)
}
