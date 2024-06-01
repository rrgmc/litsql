package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/chain"
)

type FromChain[T any] struct {
	sq.ModTagImpl[T]
	*iclause.From
}

func (f *FromChain[T]) Apply(a litsql.QueryBuilder) {
	a.Add(f.From)
}

func (f *FromChain[T]) As(alias string, columns ...string) chain.From[T] {
	f.SetAs(alias, columns...)
	return f
}

func (f *FromChain[T]) Only() chain.From[T] {
	f.SetOnly()
	return f
}

func (f *FromChain[T]) Lateral() chain.From[T] {
	f.SetLateral()
	return f
}

func (f *FromChain[T]) WithOrdinality() chain.From[T] {
	f.SetWithOrdinality()
	return f
}

func (f *FromChain[T]) WrappedQueryClause() litsql.QueryClause {
	return f.From
}
