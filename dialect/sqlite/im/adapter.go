package im

import (
	litsql "github.com/rrgmc/litsql"
	tag "github.com/rrgmc/litsql/dialect/sqlite/tag"
	sq "github.com/rrgmc/litsql/sq"
	chain "github.com/rrgmc/litsql/sq/chain"
)

type withChainAdapter struct {
	sq.ModTagImpl[tag.InsertTag]
	chain chain.With[tag.InsertTag]
}

func (a *withChainAdapter) Apply(apply litsql.QueryBuilder) {
	a.chain.Apply(apply)
}

func (a *withChainAdapter) As(q litsql.Query) WithChain {
	_ = a.chain.As(q)
	return a
}

func (a *withChainAdapter) Materialized() WithChain {
	_ = a.chain.Materialized()
	return a
}

func (a *withChainAdapter) NotMaterialized() WithChain {
	_ = a.chain.NotMaterialized()
	return a
}
