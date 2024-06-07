package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/sq/clause"
)

type Join struct {
	Type string
	To   *From // the expression for the table

	// Join methods
	Natural bool
	On      []litsql.Expression
	Using   []string
}

func (c *Join) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if c.Type == "" {
		return nil, internal.NewClauseError("'JOIN' type is required")
	}
	if c.To == nil {
		return nil, internal.NewClauseError("'JOIN' table is required")
	}

	b := litsql.NewExpressBuilder(w, d, start)

	w.AddSeparator(true)
	if c.Natural {
		w.Write("NATURAL ")
	}
	w.Write(c.Type)
	w.Write(" ")
	b.Express(c.To)

	b.ExpressSlice(c.On, expr.Raw(" ON "), expr.Raw(" AND "), nil)
	for k, col := range c.Using {
		if k == 0 {
			w.Write(" USING(")
		} else {
			w.Write(", ")
		}
		w.Write(col)
		if k == len(c.Using)-1 {
			w.Write(")")
		}
	}

	return b.Result()
}

var _ litsql.QueryClauseMultiple = (*Join)(nil)

func (c *Join) ClauseID() string {
	return "2a863df4-21ec-4d9b-939c-7a2ec211c9b8"
}

func (c *Join) ClauseOrder() int {
	return clause.OrderJoin
}

func (c *Join) ClauseMultiple() {}

func (c *Join) SetNatural() {
	c.Natural = true
}

func (c *Join) SetOn(on string) {
	c.On = append(c.On, expr.String(on))
}

func (c *Join) SetOnExpr(on litsql.Expression) {
	c.On = append(c.On, on)
}

func (c *Join) SetOnClause(query string, args ...any) {
	c.On = append(c.On, expr.Clause(query, args...))
}

func (c *Join) SetUsing(using ...string) {
	c.Using = append(c.Using, using...)
}
