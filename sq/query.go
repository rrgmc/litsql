package sq

import "github.com/rrgmc/litsql"

type Query[T any] interface {
	litsql.QueryBuilder
	litsql.Query
	QueryModApply[T]
	BuildQuery
}

type BuildQuery interface {
	Build(writerOptions ...WriterOption) (string, Args, error)
}
