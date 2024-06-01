package ium

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/clause"
)

func Update[T any](dialect litsql.Dialect, mods ...sq.QueryMod[T]) sq.Query[T] {
	ret := &UpdateQuery[T]{
		DefaultQuery: isq.NewDefaultQuery[T](dialect, clause.OrderUpdateStart, "UPDATE "),
	}
	for _, m := range mods {
		m.Apply(ret)
	}
	return ret
}

type UpdateQuery[T any] struct {
	*isq.DefaultQuery[T]
}
