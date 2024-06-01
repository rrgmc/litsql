package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type Returning struct {
	Expressions []litsql.Expression
}

func (c *Returning) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	w.AddSeparator(true)
	if len(c.Expressions) == 0 {
		w.Write("RETURNING *")

	}
	return litsql.ExpressSlice(w, d, start, c.Expressions, expr.Raw("RETURNING "), expr.CommaSpace, nil)
}

var _ litsql.QueryClauseMerge = (*Returning)(nil)

func (c *Returning) ClauseOrder() int {
	return clause.OrderReturning
}

func (c *Returning) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*Returning)
	if !ok {
		panic("invalid merge")
	}
	c.Expressions = append(c.Expressions, o.Expressions...)
}
