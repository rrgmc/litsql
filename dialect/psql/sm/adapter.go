package sm

import (
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/ichain"
)

// type groupByChainAdapter struct {
// 	// sq.ModTagImpl[tag.SelectTag]
// 	*ichain.GroupByChain[tag.SelectTag, GroupByChain]
// }

// type groupByChainAdapter = ichain.GroupByChain[tag.SelectTag, GroupByChain]

// func (a *groupByChainAdapter) Apply(apply litsql.QueryBuilder) {
// 	a.chain.Apply(apply)
// }
//
// func (a *groupByChainAdapter) Distinct() GroupByChain {
// 	_ = a.chain.Distinct()
// 	return a
// }

func newGroupByChain(chain *ichain.GroupByChain[tag.SelectTag, GroupByChain]) GroupByChain {
	chain.SetChainSelf(chain)
	return chain
}