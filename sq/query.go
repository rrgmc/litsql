package sq

import "github.com/rrgmc/litsql"

type Query[T any] interface {
	litsql.QueryBuilder
	litsql.Query
	QueryModApply[T]
	litsql.BuildQuery
}
