package chain

import "github.com/rrgmc/litsql/sq"

type From[T any] interface {
	sq.QueryMod[T]
	As(alias string, columns ...string) From[T]
	Only() From[T]
	Lateral() From[T]
	WithOrdinality() From[T]
}
