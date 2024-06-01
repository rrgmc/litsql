package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type OrderBy struct {
	Expressions []litsql.Expression
}

func (c *OrderBy) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if len(c.Expressions) == 0 {
		return nil, nil
	}
	w.AddSeparator(true)
	return litsql.ExpressSlice(w, d, start, c.Expressions, expr.Raw("ORDER BY "), expr.CommaSpace, nil)
}

var _ litsql.QueryClauseMerge = (*OrderBy)(nil)

func (c *OrderBy) ClauseID() string {
	return "2a543fbd-6d9e-4470-b713-774e3117eb11"
}

func (c *OrderBy) ClauseOrder() int {
	return clause.OrderOrderBy
}

func (c *OrderBy) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*OrderBy)
	if !ok {
		panic("invalid merge")
	}
	c.Expressions = append(c.Expressions, o.Expressions...)
}
