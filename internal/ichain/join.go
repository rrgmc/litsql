package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

const (
	JoinInnerJoin    = "INNER JOIN"
	JoinLeftJoin     = "LEFT JOIN"
	JoinRightJoin    = "RIGHT JOIN"
	JoinFullJoin     = "FULL JOIN"
	JoinCrossJoin    = "CROSS JOIN"
	JoinStraightJoin = "STRAIGHT_JOIN"
)

type Join[T any] interface {
	sq.QueryMod[T]
	As(alias string, columns ...string) Join[T]
	Only() Join[T]
	Lateral() Join[T]
	WithOrdinality() Join[T]
	Natural() Join[T]
	On(on string) Join[T]
	OnExpr(on litsql.Expression) Join[T]
	OnClause(query string, args ...any) Join[T]
	Using(using ...string) Join[T]
}

func NewJoinChain[T, CHAIN any](chain *JoinChain[T, CHAIN]) *JoinChain[T, CHAIN] {
	chain.Self = chain
	return chain
}

type JoinChain[T, CHAIN any] struct {
	sq.ModTagImpl[T]
	*iclause.Join
	Self any
}

func (c *JoinChain[T, CHAIN]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(c.Join)
}

func (c *JoinChain[T, CHAIN]) As(alias string, columns ...string) CHAIN {
	c.To.SetAs(alias, columns...)
	return c.Self.(CHAIN)
}

func (c *JoinChain[T, CHAIN]) Only() CHAIN {
	c.To.SetOnly()
	return c.Self.(CHAIN)
}

func (c *JoinChain[T, CHAIN]) Lateral() CHAIN {
	c.To.SetLateral()
	return c.Self.(CHAIN)
}

func (c *JoinChain[T, CHAIN]) WithOrdinality() CHAIN {
	c.To.SetWithOrdinality()
	return c.Self.(CHAIN)
}

func (c *JoinChain[T, CHAIN]) Natural() CHAIN {
	c.SetNatural()
	return c.Self.(CHAIN)
}

func (c *JoinChain[T, CHAIN]) On(on string) CHAIN {
	c.SetOn(on)
	return c.Self.(CHAIN)
}

func (c *JoinChain[T, CHAIN]) OnExpr(on litsql.Expression) CHAIN {
	c.SetOnExpr(on)
	return c.Self.(CHAIN)
}

func (c *JoinChain[T, CHAIN]) OnClause(query string, args ...any) CHAIN {
	c.SetOnClause(query, args...)
	return c.Self.(CHAIN)
}

func (c *JoinChain[T, CHAIN]) Using(using ...string) CHAIN {
	c.SetUsing(using...)
	return c.Self.(CHAIN)
}
