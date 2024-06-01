package im

import (
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/imod"
	"github.com/rrgmc/litsql/sq/mod"
)

type InsertConflictUpdateMod = mod.InsertConflictUpdateMod[tag.InsertTag, imod.InsertConflictUpdateModUM]
