package im

import (
	"github.com/rrgmc/litsql/dialect/mysql/tag"
	"github.com/rrgmc/litsql/internal/imod"
	"github.com/rrgmc/litsql/sq/chain"
)

type WithChain = chain.With[tag.InsertTag]

type InsertConflictChain = chain.InsertConflictUpdate[tag.InsertTag, imod.InsertConflictUpdateModTag]
