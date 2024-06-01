package mysql

import (
	"github.com/rrgmc/litsql/dialect/mysql/tag"
	"github.com/rrgmc/litsql/internal/iim"
)

func Insert(mods ...InsertMod) InsertQuery {
	return iim.Insert[tag.InsertTag](Dialect, mods...)
}
