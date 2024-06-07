package psql

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/psql/tag"
	"github.com/rrgmc/litsql/internal/idm"
)

func Delete(mods ...DeleteMod) DeleteQuery {
	return idm.Delete[tag.DeleteTag](Dialect, mods...)
}

func DeleteRaw(rawQuery string, args ...any) DeleteQuery {
	return Delete(idm.RawQuery[tag.DeleteTag](rawQuery, args...))
}

func DeleteRawExpr(e litsql.Expression) DeleteQuery {
	return Delete(idm.RawQueryExpr[tag.DeleteTag](e))
}
