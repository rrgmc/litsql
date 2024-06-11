// Code generated by "litsql-dialectgen"; DO NOT EDIT.
package dm

import (
	litsql "github.com/rrgmc/litsql"
	sqlite "github.com/rrgmc/litsql/dialect/sqlite"
	tag "github.com/rrgmc/litsql/dialect/sqlite/tag"
	sq "github.com/rrgmc/litsql/sq"
)

type FromChain interface {
	sq.QueryMod[tag.DeleteTag]
	As(alias string, columns ...string) FromChain
}

type JoinChain interface {
	sq.QueryMod[tag.DeleteTag]
	As(alias string, columns ...string) JoinChain
	Lateral() JoinChain
	Natural() sqlite.DeleteMod
	On(on string) JoinChain
	OnClause(query string, args ...any) JoinChain
	OnExpr(on litsql.Expression) JoinChain
	Using(using ...string) JoinChain
}

type WithChain interface {
	sq.QueryMod[tag.DeleteTag]
	As(q litsql.Query) WithChain
	Materialized() WithChain
	NotMaterialized() WithChain
}
