package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type Having struct {
	Conditions []litsql.Expression
}

func (c *Having) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if len(c.Conditions) == 0 {
		return nil, nil
	}
	w.AddSeparator(true)
	return litsql.ExpressSlice(w, d, start, c.Conditions, expr.Raw("HAVING "), expr.Raw(" AND "), nil)
}

var _ litsql.QueryClauseMerge = (*Having)(nil)

func (c *Having) ClauseOrder() int {
	return clause.OrderHaving
}

func (c *Having) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*Having)
	if !ok {
		panic("invalid merge")
	}
	c.Conditions = append(c.Conditions, o.Conditions...)
}
