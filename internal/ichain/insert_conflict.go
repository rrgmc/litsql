package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/imod"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/mod"
)

type InsertConflictUpdate[T, UM any] interface {
	sq.QueryMod[T]
	Where(condition string) InsertConflictUpdate[T, UM]
	WhereExpr(condition litsql.Expression) InsertConflictUpdate[T, UM]
	WhereClause(query string, args ...any) InsertConflictUpdate[T, UM]
	DoNothing() InsertConflictUpdate[T, UM]
	DoUpdate(mods ...mod.InsertConflictUpdateMod[T, UM]) InsertConflictUpdate[T, UM]
}

func NewInsertConflictUpdateChain[T, CHAIN any](chain *InsertConflictUpdateChain[T, CHAIN]) *InsertConflictUpdateChain[T, CHAIN] {
	chain.Self = chain
	return chain
}

type InsertConflictUpdateChain[T, CHAIN any] struct {
	sq.ModTagImpl[T]
	*iclause.InsertConflictUpdate
	Self any
}

func (f *InsertConflictUpdateChain[T, CHAIN]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(f.InsertConflictUpdate)
}

func (f *InsertConflictUpdateChain[T, CHAIN]) Where(condition string) CHAIN {
	f.SetWhere(expr.String(condition))
	return f.Self.(CHAIN)
}

func (f *InsertConflictUpdateChain[T, CHAIN]) WhereExpr(condition litsql.Expression) CHAIN {
	f.SetWhere(condition)
	return f.Self.(CHAIN)
}

func (f *InsertConflictUpdateChain[T, CHAIN]) WhereClause(query string, args ...any) CHAIN {
	f.SetWhere(expr.Clause(query, args...))
	return f.Self.(CHAIN)
}

func (f *InsertConflictUpdateChain[T, CHAIN]) DoNothing() CHAIN {
	f.SetDoNothing()
	return f.Self.(CHAIN)
}

func (f *InsertConflictUpdateChain[T, CHAIN]) DoUpdate(mods ...mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag]) CHAIN {
	f.SetDoUpdate()
	for _, m := range mods {
		m.Apply(f.InsertConflictUpdate)
	}
	return f.Self.(CHAIN)
}
