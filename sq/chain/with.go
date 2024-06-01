package chain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

type With[T any] interface {
	sq.QueryMod[T]
	Recursive() With[T]
	As(q litsql.Query) With[T]
	NotMaterialized() With[T]
	Materialized() With[T]
	SearchBreadth(setCol string, searchCols ...string) With[T]
	SearchDepth(setCol string, searchCols ...string) With[T]
	Cycle(set, using string, cols ...string) With[T]
	CycleValue(value, defaultVal any) With[T]
}
