package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type Offset struct {
	Count litsql.Expression
}

func (c *Offset) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if c.Count == nil {
		return nil, nil
	}
	w.AddSeparator(true)
	return litsql.ExpressIf(w, d, start, c.Count, c.Count != nil, expr.Raw("OFFSET "), nil)
}

var _ litsql.QueryClause = (*Offset)(nil)

func (c *Offset) ClauseID() string {
	return "4a893d65-4f97-435c-8f5d-cb34021afb41"
}

func (c *Offset) ClauseOrder() int {
	return clause.OrderOffset
}
