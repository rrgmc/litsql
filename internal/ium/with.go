package ium

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq/chain"
)

func With[T any](name string, columns ...string) chain.With[T] {
	return WithE[T](name, expr.StringList(columns)...)
}

func WithE[T any](name string, columns ...litsql.Expression) chain.With[T] {
	cte := &iclause.CTE{
		Name:    name,
		Columns: columns,
	}
	return &ichain.WithChain[T]{
		With: &iclause.With{
			CTEs: []*iclause.CTE{cte},
		},
		CTE: cte,
	}
}
