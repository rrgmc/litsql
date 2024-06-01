package sq

import "github.com/rrgmc/litsql"

type QueryMod[T any] interface {
	ModTag[T]
	Apply(qb litsql.QueryBuilder)
}

type Mod[T, A any] interface {
	ModTag[T]
	Apply(A)
}
