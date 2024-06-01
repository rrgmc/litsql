package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type Where struct {
	Conditions []litsql.Expression
}

func (c *Where) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if len(c.Conditions) == 0 {
		return nil, nil
	}
	w.AddSeparator(true)
	return litsql.ExpressSlice(w, d, start, c.Conditions, expr.Raw("WHERE "), expr.Raw(" AND "), nil)
}

var _ litsql.QueryClauseMerge = (*Where)(nil)

func (c *Where) ClauseOrder() int {
	return clause.OrderWhere
}

func (c *Where) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*Where)
	if !ok {
		panic("invalid merge")
	}
	c.Conditions = append(c.Conditions, o.Conditions...)
}
