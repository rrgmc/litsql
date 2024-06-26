// Code generated by "litsql-dialectgen"; DO NOT EDIT.
package sm

import (
	litsql "github.com/rrgmc/litsql"
	tag "github.com/rrgmc/litsql/dialect/mysql/tag"
	ichain "github.com/rrgmc/litsql/internal/ichain"
	sq "github.com/rrgmc/litsql/sq"
)

type FromChain interface {
	sq.QueryMod[tag.SelectTag]
	As(alias string, columns ...string) FromChain
	Lateral() FromChain
}

type GroupByChain interface {
	sq.QueryMod[tag.SelectTag]
	With(with string) GroupByChain
}

type JoinChain interface {
	sq.QueryMod[tag.SelectTag]
	As(alias string, columns ...string) JoinChain
	Lateral() JoinChain
	Natural() JoinChain
	On(on string) JoinChain
	OnExpr(on litsql.Expression) JoinChain
	OnClause(query string, args ...any) JoinChain
	Using(using ...string) JoinChain
}

type WindowChain interface {
	sq.QueryMod[tag.SelectTag]
	From(name string) WindowChain
	PartitionBy(condition ...string) WindowChain
	PartitionByExpr(condition ...litsql.Expression) WindowChain
	OrderBy(order ...string) WindowChain
	OrderByExpr(order ...litsql.Expression) WindowChain
	Frame(frame string) WindowChain
	FrameExpr(frame litsql.Expression) WindowChain
	FrameClause(query string, args ...any) WindowChain
}

type WithChain interface {
	sq.QueryMod[tag.SelectTag]
	Recursive() WithChain
	As(q litsql.Query) WithChain
}

// ensure interface is implemented by source type

var _ FromChain = (*ichain.FromChain[tag.SelectTag, FromChain])(nil)

var _ GroupByChain = (*ichain.GroupByChain[tag.SelectTag, GroupByChain])(nil)

var _ JoinChain = (*ichain.JoinChain[tag.SelectTag, JoinChain])(nil)

var _ WindowChain = (*ichain.WindowChain[tag.SelectTag, WindowChain])(nil)

var _ WithChain = (*ichain.WithChain[tag.SelectTag, WithChain])(nil)
