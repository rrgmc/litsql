package chain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/mod"
)

type InsertConflict[T, UM any] interface {
	sq.QueryMod[T]
	Where(condition string) InsertConflict[T, UM]
	WhereE(condition litsql.Expression) InsertConflict[T, UM]
	WhereC(query string, args ...any) InsertConflict[T, UM]
	DoNothing() InsertConflict[T, UM]
	DoUpdate(mods ...mod.InsertConflictUpdateMod[T, UM]) InsertConflict[T, UM]
}
