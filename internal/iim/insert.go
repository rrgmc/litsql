package iim

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/clause"
)

func Insert[T any](dialect litsql.Dialect, mods ...sq.QueryMod[T]) sq.Query[T] {
	ret := &InsertQuery[T]{
		DefaultQuery: isq.NewDefaultQuery[T](dialect, clause.OrderInsertStart, "INSERT INTO "),
	}
	for _, m := range mods {
		m.Apply(ret)
	}
	return ret
}

type InsertQuery[T any] struct {
	*isq.DefaultQuery[T]
}
