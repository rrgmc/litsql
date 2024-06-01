package mysql

import (
	"github.com/rrgmc/litsql/dialect/mysql/tag"
	"github.com/rrgmc/litsql/internal/ium"
)

func Update(mods ...UpdateMod) UpdateQuery {
	return ium.Update[tag.UpdateTag](Dialect, mods...)
}
