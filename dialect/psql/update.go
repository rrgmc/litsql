package psql

import (
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/ium"
)

func Update(mods ...UpdateMod) UpdateQuery {
	return ium.Update[tag.UpdateTag](Dialect, mods...)
}
