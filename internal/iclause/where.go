package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
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

func (c *Where) ClauseID() string {
	return "b3298776-3415-41a9-b8df-fa65a838fb47"
}

func (c *Where) ClauseOrder() int {
	return clause.OrderWhere
}

func (c *Where) ClauseMerge(other litsql.QueryClause) error {
	o, ok := other.(*Where)
	if !ok {
		return internal.NewClauseErrorInvalidMerge("Where")
	}
	c.Conditions = append(c.Conditions, o.Conditions...)
	return nil
}
