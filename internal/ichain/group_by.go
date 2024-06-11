package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

type GroupBy[T, SELF any] interface {
	sq.QueryMod[T]
	Distinct() SELF
	With(with string) SELF
}

type GroupByChain[T, SELF any] struct {
	sq.ModTagImpl[T]
	self SELF
	*iclause.GroupBy
}

func (c *GroupByChain[T, SELF]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(c.GroupBy)
}

func (c *GroupByChain[T, SELF]) Distinct() SELF {
	c.SetGroupByDistinct(true)
	return c.self
}

func (c *GroupByChain[T, SELF]) With(with string) SELF {
	c.SetGroupWith(with)
	return c.self
}
