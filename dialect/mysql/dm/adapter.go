package dm

import (
	litsql "github.com/rrgmc/litsql"
	mysql "github.com/rrgmc/litsql/dialect/mysql"
	tag "github.com/rrgmc/litsql/dialect/mysql/tag"
	sq "github.com/rrgmc/litsql/sq"
	chain "github.com/rrgmc/litsql/sq/chain"
)

type fromChainAdapter struct {
	sq.ModTagImpl[tag.DeleteTag]
	chain chain.From[tag.DeleteTag]
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

type joinChainAdapter struct {
	sq.ModTagImpl[tag.DeleteTag]
	chain chain.Join[tag.DeleteTag]
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

func (a *joinChainAdapter) Natural() mysql.DeleteMod {
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
	sq.ModTagImpl[tag.DeleteTag]
	chain chain.With[tag.DeleteTag]
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
