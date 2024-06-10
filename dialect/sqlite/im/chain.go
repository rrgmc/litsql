package im

import (
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/internal/imod"
	"github.com/rrgmc/litsql/sq/chain"
)

type InsertConflictChain = chain.InsertConflictUpdate[tag.InsertTag, imod.InsertConflictUpdateModTag]

type WithChain = chain.With[tag.InsertTag]
