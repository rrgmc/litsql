package mysql

import (
	"github.com/rrgmc/litsql/dialect/mysql/tag"
	"github.com/rrgmc/litsql/internal/ism"
)

func Select(mods ...SelectMod) SelectQuery {
	return ism.Select[tag.SelectTag](Dialect, mods...)
}
