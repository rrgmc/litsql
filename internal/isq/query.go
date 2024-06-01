package isq

import "github.com/rrgmc/litsql"

type Query[T any] interface {
	litsql.Query
}
