// Code generated by "litsql-dialectgen"; DO NOT EDIT.
package im

import (
	litsql "github.com/rrgmc/litsql"
	tag "github.com/rrgmc/litsql/dialect/psql/tag"
	ichain "github.com/rrgmc/litsql/internal/ichain"
	sq "github.com/rrgmc/litsql/sq"
)

type InsertConflictUpdateChain interface {
	sq.QueryMod[tag.InsertTag]
	Where(condition string) InsertConflictUpdateChain
	WhereExpr(condition litsql.Expression) InsertConflictUpdateChain
	WhereClause(query string, args ...any) InsertConflictUpdateChain
	DoNothing() InsertConflictUpdateChain
	DoUpdate(mods ...InsertConflictUpdateMod) InsertConflictUpdateChain
}

type WithChain interface {
	sq.QueryMod[tag.InsertTag]
	Recursive() WithChain
	As(q litsql.Query) WithChain
	NotMaterialized() WithChain
	Materialized() WithChain
	SearchBreadth(setCol string, searchCols ...string) WithChain
	SearchDepth(setCol string, searchCols ...string) WithChain
	Cycle(set string, using string, cols ...string) WithChain
	CycleValue(value any, defaultVal any) WithChain
}

// ensure interface is implemented by source type

var _ InsertConflictUpdateChain = (*ichain.InsertConflictUpdateChain[tag.InsertTag, InsertConflictUpdateChain])(nil)

var _ WithChain = (*ichain.WithChain[tag.InsertTag, WithChain])(nil)
