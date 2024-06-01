package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type Columns struct {
	Columns []litsql.Expression
}

func (c *Columns) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, d, start)
	if len(c.Columns) > 0 {
		b.ExpressSlice(c.Columns, nil, expr.Raw(", "), nil)
	} else {
		w.Write("*")
	}
	return b.Result()
}

var _ litsql.QueryClauseMerge = (*Columns)(nil)

func (c *Columns) ClauseID() string {
	return "e0134442-a30a-462c-a885-e21b033c8e21"
}

func (c *Columns) ClauseOrder() int {
	return clause.OrderColumns
}

func (c *Columns) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*Columns)
	if !ok {
		panic("invalid merge")
	}
	c.Columns = append(c.Columns, o.Columns...)
}
