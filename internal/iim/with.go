package iim

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/iclause"
)

func With[T, CHAIN any](name string, columns ...string) *ichain.WithChain[T, CHAIN] {
	return WithExpr[T, CHAIN](name, expr.StringList(columns)...)
}

func WithExpr[T, CHAIN any](name string, columns ...litsql.Expression) *ichain.WithChain[T, CHAIN] {
	cte := &iclause.CTE{
		Name:    name,
		Columns: columns,
	}
	return ichain.NewWithChain[T, CHAIN](&ichain.WithChain[T, CHAIN]{
		With: &iclause.With{
			CTEs: []*iclause.CTE{cte},
		},
		CTE: cte,
	})
}
