package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/imod"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/chain"
	"github.com/rrgmc/litsql/sq/mod"
)

type InsertConflictChain[T any] struct {
	sq.ModTagImpl[T]
	*iclause.InsertConflict
}

func (f *InsertConflictChain[T]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(f.InsertConflict)
}

func (f *InsertConflictChain[T]) Where(condition string) chain.InsertConflict[T, imod.InsertConflictUpdateModUM] {
	f.SetWhere(expr.S(condition))
	return f
}

func (f *InsertConflictChain[T]) WhereE(condition litsql.Expression) chain.InsertConflict[T, imod.InsertConflictUpdateModUM] {
	f.SetWhere(condition)
	return f
}

func (f *InsertConflictChain[T]) WhereC(query string, args ...any) chain.InsertConflict[T, imod.InsertConflictUpdateModUM] {
	f.SetWhere(expr.C(query, args...))
	return f
}

func (f *InsertConflictChain[T]) DoNothing() chain.InsertConflict[T, imod.InsertConflictUpdateModUM] {
	f.SetDoNothing()
	return f
}

func (f *InsertConflictChain[T]) DoUpdate(mods ...mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModUM]) chain.InsertConflict[T, imod.InsertConflictUpdateModUM] {
	f.SetDoUpdate()
	for _, m := range mods {
		m.Apply(f.InsertConflict)
	}
	return f
}
