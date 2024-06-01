package chain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

type Join[T any] interface {
	sq.QueryMod[T]
	As(alias string, columns ...string) Join[T]
	Only() Join[T]
	Lateral() Join[T]
	WithOrdinality() Join[T]
	Natural() sq.QueryMod[T]
	On(on string) Join[T]
	OnE(on litsql.Expression) Join[T]
	OnC(query string, args ...any) Join[T]
	Using(using ...string) Join[T]
}
