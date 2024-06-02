package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/chain"
)

const (
	JoinInnerJoin    = "INNER JOIN"
	JoinLeftJoin     = "LEFT JOIN"
	JoinRightJoin    = "RIGHT JOIN"
	JoinFullJoin     = "FULL JOIN"
	JoinCrossJoin    = "CROSS JOIN"
	JoinStraightJoin = "STRAIGHT_JOIN"
)

type JoinChain[T any] struct {
	sq.ModTagImpl[T]
	*iclause.Join
}

func (c *JoinChain[T]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(c.Join)
}

func (c *JoinChain[T]) As(alias string, columns ...string) chain.Join[T] {
	c.To.SetAs(alias, columns...)
	return c
}

func (c *JoinChain[T]) Only() chain.Join[T] {
	c.To.SetOnly()
	return c
}

func (c *JoinChain[T]) Lateral() chain.Join[T] {
	c.To.SetLateral()
	return c
}

func (c *JoinChain[T]) WithOrdinality() chain.Join[T] {
	c.To.SetWithOrdinality()
	return c
}

func (c *JoinChain[T]) Natural() sq.QueryMod[T] {
	c.SetNatural()
	return c
}

func (c *JoinChain[T]) On(on string) chain.Join[T] {
	c.SetOn(on)
	return c
}

func (c *JoinChain[T]) OnE(on litsql.Expression) chain.Join[T] {
	c.SetOnE(on)
	return c
}

func (c *JoinChain[T]) OnC(query string, args ...any) chain.Join[T] {
	c.SetOnC(query, args...)
	return c
}

func (c *JoinChain[T]) Using(using ...string) chain.Join[T] {
	c.SetUsing(using...)
	return c
}
