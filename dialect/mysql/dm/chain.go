package dm

import (
	"github.com/rrgmc/litsql/dialect/mysql/tag"
	"github.com/rrgmc/litsql/sq/chain"
)

type FromChain = chain.From[tag.DeleteTag]

type JoinChain = chain.Join[tag.DeleteTag]

type WithChain = chain.With[tag.DeleteTag]
