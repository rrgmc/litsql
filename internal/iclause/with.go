package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type With struct {
	Recursive bool
	CTEs      []*CTE
}

func (c *With) WriteSQL(wr litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	wr.AddSeparator(true)
	prefix := "WITH "
	if c.Recursive {
		prefix = "WITH RECURSIVE "
	}
	return litsql.ExpressSlice(wr, d, start, expr.CastSlice(c.CTEs), expr.Raw(prefix), expr.CommaWriterSeparator, nil)
}

var _ litsql.QueryClauseMerge = (*With)(nil)

func (c *With) ClauseID() string {
	return "4f3db589-e12c-412c-8af5-5f2d1eb9778e"
}

func (c *With) ClauseOrder() int {
	return clause.OrderWith
}

func (c *With) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*With)
	if !ok {
		panic("invalid merge")
	}
	if o.Recursive {
		c.Recursive = o.Recursive
	}
	c.CTEs = append(c.CTEs, o.CTEs...)
}

func (c *With) SetRecursive(r bool) {
	c.Recursive = r
}
