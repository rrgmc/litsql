package idm

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/clause"
)

func Delete[T any](dialect litsql.Dialect, mods ...sq.QueryMod[T]) sq.Query[T] {
	ret := &DeleteQuery[T]{
		DefaultQuery: isq.NewDefaultQuery[T](dialect, clause.OrderDeleteStart, "DELETE FROM "),
	}
	for _, m := range mods {
		m.Apply(ret)
	}
	return ret
}

type DeleteQuery[T any] struct {
	*isq.DefaultQuery[T]
}
