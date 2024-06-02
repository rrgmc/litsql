package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/chain"
)

type WithChain[T any] struct {
	sq.ModTagImpl[T]
	*iclause.With
	CTE *iclause.CTE
}

func (c *WithChain[T]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(c.With)
}

func (c *WithChain[T]) Recursive() chain.With[T] {
	c.SetRecursive(true)
	return c
}

func (c *WithChain[T]) As(q litsql.Query) chain.With[T] {
	c.CTE.SetQuery(q)
	return c
}

func (c *WithChain[T]) NotMaterialized() chain.With[T] {
	c.CTE.SetNotMaterialized()
	return c
}

func (c *WithChain[T]) Materialized() chain.With[T] {
	c.CTE.SetMaterialized()
	return c
}

func (c *WithChain[T]) SearchBreadth(setCol string, searchCols ...string) chain.With[T] {
	c.CTE.Search = iclause.CTESearch{
		Order:   iclause.SearchBreadth,
		Columns: searchCols,
		Set:     setCol,
	}
	return c
}

func (c *WithChain[T]) SearchDepth(setCol string, searchCols ...string) chain.With[T] {
	c.CTE.Search = iclause.CTESearch{
		Order:   iclause.SearchDepth,
		Columns: searchCols,
		Set:     setCol,
	}
	return c
}

func (c *WithChain[T]) Cycle(set, using string, cols ...string) chain.With[T] {
	c.CTE.Cycle.Set = set
	c.CTE.Cycle.Using = using
	c.CTE.Cycle.Columns = cols
	return c
}

func (c *WithChain[T]) CycleValue(value, defaultVal any) chain.With[T] {
	c.CTE.Cycle.SetVal = expr.Arg(value)
	c.CTE.Cycle.DefaultVal = expr.Arg(defaultVal)
	return c
}
