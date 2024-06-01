package sqlite

import (
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/internal/idm"
)

func Delete(mods ...DeleteMod) DeleteQuery {
	return idm.Delete[tag.DeleteTag](Dialect, mods...)
}
