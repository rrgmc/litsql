package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/imod"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/mod"
)

type InsertConflict[T, UM any] interface {
	sq.QueryMod[T]
	Where(condition string) InsertConflict[T, UM]
	WhereExpr(condition litsql.Expression) InsertConflict[T, UM]
	WhereClause(query string, args ...any) InsertConflict[T, UM]
	DoNothing() InsertConflict[T, UM]
	DoUpdate(mods ...mod.InsertConflictUpdateMod[T, UM]) InsertConflict[T, UM]
}

func NewInsertConflictChain[T, CHAIN any](chain *InsertConflictChain[T, CHAIN]) *InsertConflictChain[T, CHAIN] {
	chain.Self = chain
	return chain
}

type InsertConflictChain[T, CHAIN any] struct {
	sq.ModTagImpl[T]
	*iclause.InsertConflictUpdate
	Self any
}

func (f *InsertConflictChain[T, CHAIN]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(f.InsertConflictUpdate)
}

func (f *InsertConflictChain[T, CHAIN]) Where(condition string) CHAIN {
	f.SetWhere(expr.String(condition))
	return f.Self.(CHAIN)
}

func (f *InsertConflictChain[T, CHAIN]) WhereExpr(condition litsql.Expression) CHAIN {
	f.SetWhere(condition)
	return f.Self.(CHAIN)
}

func (f *InsertConflictChain[T, CHAIN]) WhereClause(query string, args ...any) CHAIN {
	f.SetWhere(expr.Clause(query, args...))
	return f.Self.(CHAIN)
}

func (f *InsertConflictChain[T, CHAIN]) DoNothing() CHAIN {
	f.SetDoNothing()
	return f.Self.(CHAIN)
}

func (f *InsertConflictChain[T, CHAIN]) DoUpdate(mods ...mod.InsertConflictUpdateMod[T, imod.InsertConflictUpdateModTag]) CHAIN {
	f.SetDoUpdate()
	for _, m := range mods {
		m.Apply(f.InsertConflictUpdate)
	}
	return f.Self.(CHAIN)
}
