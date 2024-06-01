package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq/chain"
)

func GroupBy[T any](columns ...string) chain.GroupBy[T] {
	return GroupByE[T](expr.SL(columns)...)
}

func GroupByE[T any](columns ...litsql.Expression) chain.GroupBy[T] {
	return &ichain.GroupByChain[T]{
		GroupBy: &iclause.GroupBy{
			Groups: columns,
		},
	}
}
