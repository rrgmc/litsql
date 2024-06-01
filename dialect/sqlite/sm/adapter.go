package sm

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/sqlite/tag"
	"github.com/rrgmc/litsql/sq"
	"github.com/rrgmc/litsql/sq/chain"
)

type fromChainAdapter struct {
	sq.ModTagImpl[tag.SelectTag]
	chain chain.From[tag.SelectTag]
}

func (f *fromChainAdapter) Apply(apply litsql.QueryBuilder) {
	f.chain.Apply(apply)
}

func (f *fromChainAdapter) As(alias string, columns ...string) FromChain {
	_ = f.chain.As(alias, columns...)
	return f
}
