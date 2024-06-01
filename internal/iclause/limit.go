package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type Limit struct {
	Count litsql.Expression
}

func (c *Limit) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if c.Count == nil {
		return nil, nil
	}
	w.AddSeparator(true)
	return litsql.ExpressIf(w, d, start, c.Count, c.Count != nil, expr.Raw("LIMIT "), nil)
}

var _ litsql.QueryClause = (*Limit)(nil)

func (c *Limit) ClauseOrder() int {
	return clause.OrderLimit
}
