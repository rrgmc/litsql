package sm

import (
	litsql "github.com/rrgmc/litsql"
	tag "github.com/rrgmc/litsql/dialect/psql/tag"
	sq "github.com/rrgmc/litsql/sq"
	chain "github.com/rrgmc/litsql/sq/chain"
)

type groupByChainAdapter struct {
	sq.ModTagImpl[tag.SelectTag]
	chain chain.GroupBy[tag.SelectTag]
}

func (a *groupByChainAdapter) Apply(apply litsql.QueryBuilder) {
	a.chain.Apply(apply)
}

func (a *groupByChainAdapter) Distinct() GroupByChain {
	_ = a.chain.Distinct()
	return a
}
