package sm

import (
	litsql "github.com/rrgmc/litsql"
	sqlite "github.com/rrgmc/litsql/dialect/sqlite"
	tag "github.com/rrgmc/litsql/dialect/sqlite/tag"
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

type groupByChainAdapter struct {
	sq.ModTagImpl[tag.SelectTag]
	chain chain.GroupBy[tag.SelectTag]
}

func (a *groupByChainAdapter) Apply(apply litsql.QueryBuilder) {
	a.chain.Apply(apply)
}

func (a *groupByChainAdapter) Distinct() GroupByChain {
	_ = a.chain.Distinct()
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

func (a *joinChainAdapter) Natural() sqlite.SelectMod {
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

func (a *withChainAdapter) Materialized() WithChain {
	_ = a.chain.Materialized()
	return a
}

func (a *withChainAdapter) NotMaterialized() WithChain {
	_ = a.chain.NotMaterialized()
	return a
}
