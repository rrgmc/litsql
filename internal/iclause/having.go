package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
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

func (c *Having) ClauseID() string {
	return "5d15bb66-3688-48f9-8619-134c042f4953"
}

func (c *Having) ClauseOrder() int {
	return clause.OrderHaving
}

func (c *Having) ClauseMerge(other litsql.QueryClause) error {
	o, ok := other.(*Having)
	if !ok {
		return internal.NewClauseErrorInvalidMerge("Having")
	}
	c.Conditions = append(c.Conditions, o.Conditions...)
	return nil
}
