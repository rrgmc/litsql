package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

type From[T any] interface {
	sq.QueryMod[T]
	As(alias string, columns ...string) From[T]
	Only() From[T]
	Lateral() From[T]
	WithOrdinality() From[T]
}

func NewFromChain[T, CHAIN any](chain *FromChain[T, CHAIN]) *FromChain[T, CHAIN] {
	chain.Self = chain
	return chain
}

type FromChain[T, CHAIN any] struct {
	sq.ModTagImpl[T]
	*iclause.From
	Self any
}

func (f *FromChain[T, CHAIN]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(f.From)
}

func (f *FromChain[T, CHAIN]) As(alias string, columns ...string) CHAIN {
	f.SetAs(alias, columns...)
	return f.Self.(CHAIN)
}

func (f *FromChain[T, CHAIN]) Only() CHAIN {
	f.SetOnly()
	return f.Self.(CHAIN)
}

func (f *FromChain[T, CHAIN]) Lateral() CHAIN {
	f.SetLateral()
	return f.Self.(CHAIN)
}

func (f *FromChain[T, CHAIN]) WithOrdinality() CHAIN {
	f.SetWithOrdinality()
	return f.Self.(CHAIN)
}
