package sm

import (
	litsql "github.com/rrgmc/litsql"
	mysql "github.com/rrgmc/litsql/dialect/mysql"
	tag "github.com/rrgmc/litsql/dialect/mysql/tag"
	sq "github.com/rrgmc/litsql/sq"
	chain "github.com/rrgmc/litsql/sq/chain"
)

type fromChainAdapter struct {
	sq.ModTagImpl[tag.SelectTag]
	chain chain.From[tag.SelectTag]
}

func (a *fromChainAdapter) Apply(apply litsql.QueryBuilder) {
	a.chain.Apply(apply)
}

func (a *fromChainAdapter) As(alias string, columns ...string) FromChain {
	_ = a.chain.As(alias, columns...)
	return a
}

func (a *fromChainAdapter) Lateral() FromChain {
	_ = a.chain.Lateral()
	return a
}

type groupByChainAdapter struct {
	sq.ModTagImpl[tag.SelectTag]
	chain chain.GroupBy[tag.SelectTag]
}

func (a *groupByChainAdapter) Apply(apply litsql.QueryBuilder) {
	a.chain.Apply(apply)
}

func (a *groupByChainAdapter) With(with string) GroupByChain {
	_ = a.chain.With(with)
	return a
}

type joinChainAdapter struct {
	sq.ModTagImpl[tag.SelectTag]
	chain chain.Join[tag.SelectTag]
}

func (a *joinChainAdapter) Apply(apply litsql.QueryBuilder) {
	a.chain.Apply(apply)
}

func (a *joinChainAdapter) As(alias string, columns ...string) JoinChain {
	_ = a.chain.As(alias, columns...)
	return a
}

func (a *joinChainAdapter) Lateral() JoinChain {
	_ = a.chain.Lateral()
	return a
}

func (a *joinChainAdapter) Natural() mysql.SelectMod {
	_ = a.chain.Natural()
	return a
}

func (a *joinChainAdapter) On(on string) JoinChain {
	_ = a.chain.On(on)
	return a
}

func (a *joinChainAdapter) OnClause(query string, args ...any) JoinChain {
	_ = a.chain.OnClause(query, args...)
	return a
}

func (a *joinChainAdapter) OnExpr(on litsql.Expression) JoinChain {
	_ = a.chain.OnExpr(on)
	return a
}

func (a *joinChainAdapter) Using(using ...string) JoinChain {
	_ = a.chain.Using(using...)
	return a
}

type windowChainAdapter struct {
	sq.ModTagImpl[tag.SelectTag]
	chain chain.Window[tag.SelectTag]
}

func (a *windowChainAdapter) Apply(apply litsql.QueryBuilder) {
	a.chain.Apply(apply)
}

func (a *windowChainAdapter) From(name string) WindowChain {
	_ = a.chain.From(name)
	return a
}

func (a *windowChainAdapter) FromCurrentRow() WindowChain {
	_ = a.chain.FromCurrentRow()
	return a
}

func (a *windowChainAdapter) FromFollowing(exp litsql.Expression) WindowChain {
	_ = a.chain.FromFollowing(exp)
	return a
}

func (a *windowChainAdapter) FromPreceding(exp litsql.Expression) WindowChain {
	_ = a.chain.FromPreceding(exp)
	return a
}

func (a *windowChainAdapter) FromUnboundedPreceding() WindowChain {
	_ = a.chain.FromUnboundedPreceding()
	return a
}

func (a *windowChainAdapter) OrderBy(order ...string) WindowChain {
	_ = a.chain.OrderBy(order...)
	return a
}

func (a *windowChainAdapter) OrderByExpr(order ...litsql.Expression) WindowChain {
	_ = a.chain.OrderByExpr(order...)
	return a
}

func (a *windowChainAdapter) PartitionBy(condition ...string) WindowChain {
	_ = a.chain.PartitionBy(condition...)
	return a
}

func (a *windowChainAdapter) PartitionByExpr(condition ...litsql.Expression) WindowChain {
	_ = a.chain.PartitionByExpr(condition...)
	return a
}

func (a *windowChainAdapter) Range() WindowChain {
	_ = a.chain.Range()
	return a
}

func (a *windowChainAdapter) Rows() WindowChain {
	_ = a.chain.Rows()
	return a
}

func (a *windowChainAdapter) ToCurrentRow(count int) WindowChain {
	_ = a.chain.ToCurrentRow(count)
	return a
}

func (a *windowChainAdapter) ToFollowing(exp litsql.Expression) WindowChain {
	_ = a.chain.ToFollowing(exp)
	return a
}

func (a *windowChainAdapter) ToPreceding(exp litsql.Expression) WindowChain {
	_ = a.chain.ToPreceding(exp)
	return a
}

func (a *windowChainAdapter) ToUnboundedFollowing() WindowChain {
	_ = a.chain.ToUnboundedFollowing()
	return a
}

type withChainAdapter struct {
	sq.ModTagImpl[tag.SelectTag]
	chain chain.With[tag.SelectTag]
}

func (a *withChainAdapter) Apply(apply litsql.QueryBuilder) {
	a.chain.Apply(apply)
}

func (a *withChainAdapter) As(q litsql.Query) WithChain {
	_ = a.chain.As(q)
	return a
}

func (a *withChainAdapter) Recursive() WithChain {
	_ = a.chain.Recursive()
	return a
}
