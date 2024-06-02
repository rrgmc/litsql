package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/chain"
)

type GroupByChain[T any] struct {
	sq.ModTagImpl[T]
	*iclause.GroupBy
}

func (c *GroupByChain[T]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(c.GroupBy)
}

func (c *GroupByChain[T]) Distinct() chain.GroupBy[T] {
	c.SetGroupByDistinct(true)
	return c
}

func (c *GroupByChain[T]) With(with string) chain.GroupBy[T] {
	c.SetGroupWith(with)
	return c
}
