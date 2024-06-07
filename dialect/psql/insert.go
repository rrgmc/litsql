package psql

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/iim"
	"github.com/rrgmc/litsql/internal/ism"
)

func Insert(mods ...InsertMod) InsertQuery {
	return iim.Insert[tag.InsertTag](Dialect, mods...)
}

func InsertRaw(rawQuery string, args ...any) InsertQuery {
	return Insert(ism.RawQuery[tag.InsertTag](rawQuery, args...))
}

func InsertRawExpr(e litsql.Expression) InsertQuery {
	return Insert(ism.RawQueryExpr[tag.InsertTag](e))
}
