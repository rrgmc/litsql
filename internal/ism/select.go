package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/clause"
)

func Select[T any](dialect litsql.Dialect, mods ...sq.QueryMod[T]) sq.Query[T] {
	ret := &SelectQuery[T]{
		DefaultQuery: isq.NewDefaultQuery[T](dialect, clause.OrderSelectStart, "SELECT "),
	}
	for _, m := range mods {
		m.Apply(ret)
	}
	return ret
}

type SelectQuery[T any] struct {
	*isq.DefaultQuery[T]
}
