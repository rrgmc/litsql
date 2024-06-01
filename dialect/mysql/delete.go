package mysql

import (
	"github.com/rrgmc/litsql/dialect/mysql/tag"
	"github.com/rrgmc/litsql/internal/idm"
)

func Delete(mods ...DeleteMod) DeleteQuery {
	return idm.Delete[tag.DeleteTag](Dialect, mods...)
}
