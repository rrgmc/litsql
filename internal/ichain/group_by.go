package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

type GroupBy[T, CHAIN any] interface {
	sq.QueryMod[T]
	Distinct() CHAIN
	With(with string) CHAIN
}

type GroupByChain[T, CHAIN any] struct {
	sq.ModTagImpl[T]
	self CHAIN
	*iclause.GroupBy
}

func (c *GroupByChain[T, CHAIN]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(c.GroupBy)
}

func (c *GroupByChain[T, CHAIN]) Distinct() CHAIN {
	c.SetGroupByDistinct(true)
	return c.self
}

func (c *GroupByChain[T, CHAIN]) With(with string) CHAIN {
	c.SetGroupWith(with)
	return c.self
}
