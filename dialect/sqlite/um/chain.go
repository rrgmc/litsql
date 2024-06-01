package um

import (
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/sq/chain"
)

type FromChain = chain.From[tag.UpdateTag]

type JoinChain = chain.Join[tag.UpdateTag]

type WithChain = chain.With[tag.UpdateTag]
