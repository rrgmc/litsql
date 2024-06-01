package psql

import (
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/iim"
)

func Insert(mods ...InsertMod) InsertQuery {
	return iim.Insert[tag.InsertTag](Dialect, mods...)
}
