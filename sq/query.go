package sq

import "github.com/rrgmc/litsql"

// Query is the interface for base queries.
type Query[T any] interface {
	litsql.QueryBuilder
	litsql.Query
	QueryModApply[T]
	BuildQuery
}

// BuildQuery builds query strings and arguments.
type BuildQuery interface {
	Build(options ...BuildOption) (string, []any, error)
}
