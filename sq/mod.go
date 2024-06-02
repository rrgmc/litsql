package sq

import "github.com/rrgmc/litsql"

// QueryMod is a mod for [litsql.QueryBuilder].
type QueryMod[T any] interface {
	ModTag[T]
	Apply(qb litsql.QueryBuilder)
}

// Mod is a mod for a generic type.
type Mod[T, A any] interface {
	ModTag[T]
	Apply(A)
}
