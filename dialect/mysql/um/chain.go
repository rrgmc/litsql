// Code generated by "litsql-dialectgen"; DO NOT EDIT.
package um

import (
	litsql "github.com/rrgmc/litsql"
	mysql "github.com/rrgmc/litsql/dialect/mysql"
	tag "github.com/rrgmc/litsql/dialect/mysql/tag"
	sq "github.com/rrgmc/litsql/sq"
	chain "github.com/rrgmc/litsql/sq/chain"
)

type FromChain interface {
	sq.QueryMod[tag.UpdateTag]
	As(alias string, columns ...string) FromChain
	Lateral() FromChain
}

type JoinChain interface {
	sq.QueryMod[tag.UpdateTag]
	As(alias string, columns ...string) JoinChain
	Lateral() JoinChain
	Natural() mysql.UpdateMod
	On(on string) JoinChain
	OnClause(query string, args ...any) JoinChain
	OnExpr(on litsql.Expression) JoinChain
	Using(using ...string) JoinChain
}

type WithChain = chain.With[tag.UpdateTag]
