package chain

// import (
// 	"github.com/rrgmc/litsql"
// 	"github.com/rrgmc/litsql/sq"
// 	"github.com/rrgmc/litsql/sq/mod"
// )
//
// type InsertConflictUpdate[T, UM any] interface {
// 	sq.QueryMod[T]
// 	Where(condition string) InsertConflictUpdate[T, UM]
// 	WhereExpr(condition litsql.Expression) InsertConflictUpdate[T, UM]
// 	WhereClause(query string, args ...any) InsertConflictUpdate[T, UM]
// 	DoNothing() InsertConflictUpdate[T, UM]
// 	DoUpdate(mods ...mod.InsertConflictUpdateMod[T, UM]) InsertConflictUpdate[T, UM]
// }
