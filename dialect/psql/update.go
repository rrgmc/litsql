package psql

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/ism"
	"github.com/rrgmc/litsql/internal/ium"
)

func Update(mods ...UpdateMod) UpdateQuery {
	return ium.Update[tag.UpdateTag](Dialect, mods...)
}

func UpdateRaw(rawQuery string, args ...any) UpdateQuery {
	return Update(ism.RawQuery[tag.UpdateTag](rawQuery, args...))
}

func UpdateRawExpr(e litsql.Expression) UpdateQuery {
	return Update(ism.RawQueryExpr[tag.UpdateTag](e))
}
