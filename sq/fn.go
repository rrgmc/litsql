package sq

import "github.com/rrgmc/litsql"

func QueryModFunc[T any](f func(apply litsql.QueryBuilder)) QueryMod[T] {
	return &queryModFunc[T]{f: f}
}

type queryModFunc[T any] struct {
	ModTagImpl[T]
	f func(apply litsql.QueryBuilder)
}

func (f queryModFunc[T]) Apply(a litsql.QueryBuilder) {
	f.f(a)
}
