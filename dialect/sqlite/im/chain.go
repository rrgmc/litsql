package im

import (
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/internal/imod"
	"github.com/rrgmc/litsql/sq/chain"
)

type WithChain = chain.With[tag.InsertTag]

type InsertConflictChain = chain.InsertConflict[tag.InsertTag, imod.InsertConflictUpdateModUM]
