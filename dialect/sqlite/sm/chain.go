package sm

import (
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/chain"
)

type FromChain interface {
	sq.QueryMod[tag.SelectTag]
	As(alias string, columns ...string) FromChain
}

type GroupByChain = chain.GroupBy[tag.SelectTag]

type JoinChain = chain.Join[tag.SelectTag]

type WindowChain = chain.Window[tag.SelectTag]

type WithChain = chain.With[tag.SelectTag]
