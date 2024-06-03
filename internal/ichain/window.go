package ichain

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/chain"
)

type WindowChain[T any] struct {
	sq.ModTagImpl[T]
	*iclause.Windows
	NamedWindow *iclause.NamedWindow
}

var _ chain.Window[int] = (*WindowChain[int])(nil)

func (f *WindowChain[T]) Apply(a litsql.QueryBuilder) {
	a.AddQueryClause(f.Windows)
}

func (f *WindowChain[T]) From(name string) chain.Window[T] {
	f.NamedWindow.Definition.SetFrom(name)
	return f
}

func (f *WindowChain[T]) PartitionBy(condition ...string) chain.Window[T] {
	f.NamedWindow.Definition.AddPartitionBy(expr.SL(condition)...)
	return f
}

func (f *WindowChain[T]) PartitionByE(condition ...litsql.Expression) chain.Window[T] {
	f.NamedWindow.Definition.AddPartitionBy(condition...)
	return f
}

func (f *WindowChain[T]) OrderBy(order ...string) chain.Window[T] {
	f.NamedWindow.Definition.AddOrderBy(expr.SL(order)...)
	return f
}

func (f *WindowChain[T]) OrderByE(order ...litsql.Expression) chain.Window[T] {
	f.NamedWindow.Definition.AddOrderBy(order...)
	return f
}

func (f *WindowChain[T]) Range() chain.Window[T] {
	f.NamedWindow.Definition.SetMode("RANGE")
	return f
}

func (f *WindowChain[T]) Rows() chain.Window[T] {
	f.NamedWindow.Definition.SetMode("ROWS")
	return f
}

func (f *WindowChain[T]) Groups() chain.Window[T] {
	f.NamedWindow.Definition.SetMode("GROUPS")
	return f
}

func (f *WindowChain[T]) FromUnboundedPreceding() chain.Window[T] {
	f.NamedWindow.Definition.SetStart(expr.Raw("UNBOUNDED PRECEDING"))
	return f
}

func (f *WindowChain[T]) FromPreceding(exp litsql.Expression) chain.Window[T] {
	f.NamedWindow.Definition.SetStart(litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		return litsql.ExpressIf(w, d, start, exp, true, nil, expr.Raw(" PRECEDING"))
	}))
	return f
}

func (f *WindowChain[T]) FromCurrentRow() chain.Window[T] {
	f.NamedWindow.Definition.SetStart(expr.Raw("CURRENT ROW"))
	return f
}

func (f *WindowChain[T]) FromFollowing(exp litsql.Expression) chain.Window[T] {
	f.NamedWindow.Definition.SetStart(litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		return litsql.ExpressIf(w, d, start, exp, true, nil, expr.Raw(" FOLLOWING"))
	}))
	return f
}

func (f *WindowChain[T]) ToPreceding(exp litsql.Expression) chain.Window[T] {
	f.NamedWindow.Definition.SetEnd(litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		return litsql.ExpressIf(w, d, start, exp, true, nil, expr.Raw(" PRECEDING"))
	}))
	return f
}

func (f *WindowChain[T]) ToCurrentRow(count int) chain.Window[T] {
	f.NamedWindow.Definition.SetEnd(expr.Raw("CURRENT ROW"))
	return f
}

func (f *WindowChain[T]) ToFollowing(exp litsql.Expression) chain.Window[T] {
	f.NamedWindow.Definition.SetEnd(litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		return litsql.ExpressIf(w, d, start, exp, true, nil, expr.Raw(" FOLLOWING"))
	}))
	return f
}

func (f *WindowChain[T]) ToUnboundedFollowing() chain.Window[T] {
	f.NamedWindow.Definition.SetEnd(expr.Raw("UNBOUNDED FOLLOWING"))
	return f
}

func (f *WindowChain[T]) ExcludeNoOthers() chain.Window[T] {
	f.NamedWindow.Definition.SetExclusion("NO OTHERS")
	return f
}

func (f *WindowChain[T]) ExcludeCurrentRow() chain.Window[T] {
	f.NamedWindow.Definition.SetExclusion("CURRENT ROW")
	return f
}

func (f *WindowChain[T]) ExcludeGroup() chain.Window[T] {
	f.NamedWindow.Definition.SetExclusion("GROUP")
	return f
}

func (f *WindowChain[T]) ExcludeTies() chain.Window[T] {
	f.NamedWindow.Definition.SetExclusion("TIES")
	return f
}
