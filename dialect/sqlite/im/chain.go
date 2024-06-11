// Code generated by "litsql-dialectgen"; DO NOT EDIT.
package im

import (
	litsql "github.com/rrgmc/litsql"
	tag "github.com/rrgmc/litsql/dialect/sqlite/tag"
	imod "github.com/rrgmc/litsql/internal/imod"
	sq "github.com/rrgmc/litsql/sq"
	chain "github.com/rrgmc/litsql/sq/chain"
)

type InsertConflictUpdateChain = chain.InsertConflictUpdate[tag.InsertTag, imod.InsertConflictUpdateModTag]

type WithChain interface {
	sq.QueryMod[tag.InsertTag]
	As(q litsql.Query) WithChain
	Materialized() WithChain
	NotMaterialized() WithChain
}
