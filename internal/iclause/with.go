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

func (w *With) WriteSQL(wr litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	wr.AddSeparator(true)
	prefix := "WITH "
	if w.Recursive {
		prefix = "WITH RECURSIVE "
	}
	return litsql.ExpressSlice(wr, d, start, expr.CastSlice(w.CTEs), expr.Raw(prefix), expr.CommaWriterNewLine, nil)
}

var _ litsql.QueryClauseMerge = (*With)(nil)

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

func (w *With) SetRecursive(r bool) {
	w.Recursive = r
}
