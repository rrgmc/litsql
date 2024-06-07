package psql

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/ism"
)

func Select(mods ...SelectMod) SelectQuery {
	return ism.Select[tag.SelectTag](Dialect, mods...)
}

func SelectRaw(rawQuery string, args ...any) SelectQuery {
	return Select(ism.RawQuery[tag.SelectTag](rawQuery, args...))
}

func SelectRawExpr(e litsql.Expression) SelectQuery {
	return Select(ism.RawQueryExpr[tag.SelectTag](e))
}
