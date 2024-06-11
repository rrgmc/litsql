package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

type GroupBy[T any] interface {
	sq.QueryMod[T]
	Distinct() GroupBy[T]
	With(with string) GroupBy[T]
}

func NewGroupByChain[T any](chain *GroupByChain[T, GroupBy[T]]) GroupBy[T] {
	chain.SetChainSelf(chain)
	return chain
}

type GroupByChain[T, CHAIN any] struct {
	sq.ModTagImpl[T]
	Self CHAIN
	*iclause.GroupBy
}

func (c *GroupByChain[T, CHAIN]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(c.GroupBy)
}

func (c *GroupByChain[T, CHAIN]) SetChainSelf(self CHAIN) {
	c.Self = self
}

func (c *GroupByChain[T, CHAIN]) Distinct() CHAIN {
	c.SetGroupByDistinct(true)
	return c.Self
}

func (c *GroupByChain[T, CHAIN]) With(with string) CHAIN {
	c.SetGroupWith(with)
	return c.Self
}
