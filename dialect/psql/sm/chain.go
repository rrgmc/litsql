package sm

import (
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/sq/chain"
)

type FromChain = chain.From[tag.SelectTag]

type GroupByChain = chain.GroupBy[tag.SelectTag]

type JoinChain = chain.Join[tag.SelectTag]

type WindowChain = chain.Window[tag.SelectTag]

type WithChain = chain.With[tag.SelectTag]
