package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

type With[T any] interface {
	sq.QueryMod[T]
	Recursive() With[T]
	As(q litsql.Query) With[T]
	NotMaterialized() With[T]
	Materialized() With[T]
	SearchBreadth(setCol string, searchCols ...string) With[T]
	SearchDepth(setCol string, searchCols ...string) With[T]
	Cycle(set, using string, cols ...string) With[T]
	CycleValue(value, defaultVal any) With[T]
}

func NewWithChain[T, CHAIN any](chain *WithChain[T, CHAIN]) *WithChain[T, CHAIN] {
	chain.Self = chain
	return chain
}

type WithChain[T, CHAIN any] struct {
	sq.ModTagImpl[T]
	*iclause.With
	CTE  *iclause.CTE
	Self any
}

func (c *WithChain[T, CHAIN]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(c.With)
}

func (c *WithChain[T, CHAIN]) Recursive() CHAIN {
	c.SetRecursive(true)
	return c.Self.(CHAIN)
}

func (c *WithChain[T, CHAIN]) As(q litsql.Query) CHAIN {
	c.CTE.SetQuery(q)
	return c.Self.(CHAIN)
}

func (c *WithChain[T, CHAIN]) NotMaterialized() CHAIN {
	c.CTE.SetNotMaterialized()
	return c.Self.(CHAIN)
}

func (c *WithChain[T, CHAIN]) Materialized() CHAIN {
	c.CTE.SetMaterialized()
	return c.Self.(CHAIN)
}

func (c *WithChain[T, CHAIN]) SearchBreadth(setCol string, searchCols ...string) CHAIN {
	c.CTE.Search = iclause.CTESearch{
		Order:   iclause.SearchBreadth,
		Columns: searchCols,
		Set:     setCol,
	}
	return c.Self.(CHAIN)
}

func (c *WithChain[T, CHAIN]) SearchDepth(setCol string, searchCols ...string) CHAIN {
	c.CTE.Search = iclause.CTESearch{
		Order:   iclause.SearchDepth,
		Columns: searchCols,
		Set:     setCol,
	}
	return c.Self.(CHAIN)
}

func (c *WithChain[T, CHAIN]) Cycle(set, using string, cols ...string) CHAIN {
	c.CTE.Cycle.Set = set
	c.CTE.Cycle.Using = using
	c.CTE.Cycle.Columns = cols
	return c.Self.(CHAIN)
}

func (c *WithChain[T, CHAIN]) CycleValue(value, defaultVal any) CHAIN {
	c.CTE.Cycle.SetVal = expr.Arg(value)
	c.CTE.Cycle.DefaultVal = expr.Arg(defaultVal)
	return c.Self.(CHAIN)
}
