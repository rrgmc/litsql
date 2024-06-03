package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
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

func (c *Returning) ClauseID() string {
	return "00d5282c-5e96-4dde-ba77-e120a9a8f453"
}

func (c *Returning) ClauseOrder() int {
	return clause.OrderReturning
}

func (c *Returning) ClauseMerge(other litsql.QueryClause) error {
	o, ok := other.(*Returning)
	if !ok {
		return internal.NewClauseErrorInvalidMerge("Returning")
	}
	c.Expressions = append(c.Expressions, o.Expressions...)
	return nil
}
