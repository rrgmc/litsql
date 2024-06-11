package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
)

func GroupBy[T, CHAIN any](columns ...string) *ichain.GroupByChain[T, CHAIN] {
	return GroupByExpr[T, CHAIN](expr.StringList(columns)...)
}

func GroupByExpr[T, CHAIN any](columns ...litsql.Expression) *ichain.GroupByChain[T, CHAIN] {
	return ichain.NewGroupByChain[T, CHAIN](&ichain.GroupByChain[T, CHAIN]{
		GroupBy: &iclause.GroupBy{
			Groups: columns,
		},
	})
	// return &ichain.GroupByChain[T, CHAIN]{
	// 	GroupBy: &iclause.GroupBy{
	// 		Groups: columns,
	// 	},
	// }
}
