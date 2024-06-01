package sqlite

import (
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/internal/ium"
)

func Update(mods ...UpdateMod) UpdateQuery {
	return ium.Update[tag.UpdateTag](Dialect, mods...)
}
