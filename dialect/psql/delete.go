package psql

import (
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/idm"
)

func Delete(mods ...DeleteMod) DeleteQuery {
	return idm.Delete[tag.DeleteTag](Dialect, mods...)
}
