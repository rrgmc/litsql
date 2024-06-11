package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
)

func GroupBy[T, SELF any](columns ...string) *ichain.GroupByChain[T, SELF] {
	return GroupByExpr[T, SELF](expr.StringList(columns)...)
}

func GroupByExpr[T, SELF any](columns ...litsql.Expression) *ichain.GroupByChain[T, SELF] {
	return &ichain.GroupByChain[T, SELF]{
		GroupBy: &iclause.GroupBy{
			Groups: columns,
		},
	}
}
