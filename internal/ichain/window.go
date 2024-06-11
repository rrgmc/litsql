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

func (f *WindowChain[T, CHAIN]) Range() CHAIN {
	f.NamedWindow.Definition.SetMode("RANGE")
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) Rows() CHAIN {
	f.NamedWindow.Definition.SetMode("ROWS")
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) Groups() CHAIN {
	f.NamedWindow.Definition.SetMode("GROUPS")
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) FromUnboundedPreceding() CHAIN {
	f.NamedWindow.Definition.SetStart(expr.Raw("UNBOUNDED PRECEDING"))
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) FromPreceding(exp litsql.Expression) CHAIN {
	f.NamedWindow.Definition.SetStart(litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		return litsql.ExpressIf(w, d, start, exp, true, nil, expr.Raw(" PRECEDING"))
	}))
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) FromCurrentRow() CHAIN {
	f.NamedWindow.Definition.SetStart(expr.Raw("CURRENT ROW"))
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) FromFollowing(exp litsql.Expression) CHAIN {
	f.NamedWindow.Definition.SetStart(litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		return litsql.ExpressIf(w, d, start, exp, true, nil, expr.Raw(" FOLLOWING"))
	}))
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) ToPreceding(exp litsql.Expression) CHAIN {
	f.NamedWindow.Definition.SetEnd(litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		return litsql.ExpressIf(w, d, start, exp, true, nil, expr.Raw(" PRECEDING"))
	}))
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) ToCurrentRow(count int) CHAIN {
	f.NamedWindow.Definition.SetEnd(expr.Raw("CURRENT ROW"))
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) ToFollowing(exp litsql.Expression) CHAIN {
	f.NamedWindow.Definition.SetEnd(litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		return litsql.ExpressIf(w, d, start, exp, true, nil, expr.Raw(" FOLLOWING"))
	}))
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) ToUnboundedFollowing() CHAIN {
	f.NamedWindow.Definition.SetEnd(expr.Raw("UNBOUNDED FOLLOWING"))
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) ExcludeNoOthers() CHAIN {
	f.NamedWindow.Definition.SetExclusion("NO OTHERS")
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) ExcludeCurrentRow() CHAIN {
	f.NamedWindow.Definition.SetExclusion("CURRENT ROW")
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) ExcludeGroup() CHAIN {
	f.NamedWindow.Definition.SetExclusion("GROUP")
	return f.Self.(CHAIN)
}

func (f *WindowChain[T, CHAIN]) ExcludeTies() CHAIN {
	f.NamedWindow.Definition.SetExclusion("TIES")
	return f.Self.(CHAIN)
}
