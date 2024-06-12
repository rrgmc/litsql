package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

type Window[T any] interface {
	sq.QueryMod[T]
	From(name string) Window[T]
	PartitionBy(condition ...string) Window[T]
	PartitionByExpr(condition ...litsql.Expression) Window[T]
	OrderBy(order ...string) Window[T]
	OrderByExpr(order ...litsql.Expression) Window[T]
	Frame(frame litsql.Expression) Window[T]
}

func NewWindowChain[T, CHAIN any](chain *WindowChain[T, CHAIN]) *WindowChain[T, CHAIN] {
	chain.Self = chain
	return chain
}

type WindowChain[T, CHAIN any] struct {
	sq.ModTagImpl[T]
	*iclause.Windows
	NamedWindow *iclause.NamedWindow
	Self        any
}

var _ Window[int] = (*WindowChain[int, Window[int]])(nil)

func (f *WindowChain[T, CHAIN]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(f.Windows)
}

func (f *WindowChain[T, CHAIN]) From(name string) CHAIN {
	f.NamedWindow.Definition.SetFrom(name)
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) PartitionBy(condition ...string) CHAIN {
	f.NamedWindow.Definition.AddPartitionBy(expr.StringList(condition)...)
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) PartitionByExpr(condition ...litsql.Expression) CHAIN {
	f.NamedWindow.Definition.AddPartitionBy(condition...)
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) OrderBy(order ...string) CHAIN {
	f.NamedWindow.Definition.AddOrderBy(expr.StringList(order)...)
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) OrderByExpr(order ...litsql.Expression) CHAIN {
	f.NamedWindow.Definition.AddOrderBy(order...)
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) Frame(frame litsql.Expression) CHAIN {
	f.NamedWindow.Definition.SetFrame(frame)
	return f.Self.(CHAIN)
}
