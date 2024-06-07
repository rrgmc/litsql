package chain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

type Window[T any] interface {
	sq.QueryMod[T]
	From(name string) Window[T]
	PartitionBy(condition ...string) Window[T]
	PartitionByExpr(condition ...litsql.Expression) Window[T]
	OrderBy(order ...string) Window[T]
	OrderByExpr(order ...litsql.Expression) Window[T]
	Range() Window[T]
	Rows() Window[T]
	Groups() Window[T]
	FromUnboundedPreceding() Window[T]
	FromPreceding(exp litsql.Expression) Window[T]
	FromCurrentRow() Window[T]
	FromFollowing(exp litsql.Expression) Window[T]
	ToPreceding(exp litsql.Expression) Window[T]
	ToCurrentRow(count int) Window[T]
	ToFollowing(exp litsql.Expression) Window[T]
	ToUnboundedFollowing() Window[T]
	ExcludeNoOthers() Window[T]
	ExcludeCurrentRow() Window[T]
	ExcludeGroup() Window[T]
	ExcludeTies() Window[T]
}
