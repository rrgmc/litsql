package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type Distinct struct {
	On []litsql.Expression
}

func (c *Distinct) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, d, start)
	w.Write("DISTINCT")
	if len(c.On) > 0 {
		w.Write(" ON ")
		b.ExpressSlice(c.On, expr.OpenPar, expr.CommaSpace, expr.ClosePar)
	}
	return b.Result()
}

var _ litsql.QueryClauseMerge = (*Distinct)(nil)

func (c *Distinct) ClauseID() string {
	return "e87e8c49-6dcd-4a8e-a47d-5ef8dfd3cb2a"
}

func (c *Distinct) ClauseOrder() int {
	return clause.OrderDistinct
}

func (c *Distinct) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*Distinct)
	if !ok {
		panic("invalid merge")
	}
	c.On = append(c.On, o.On...)
}
