package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type Set struct {
	Set     []litsql.Expression
	Starter bool
}

func (c *Set) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if c.Starter {
		w.AddSeparator(true)
		w.Write("SET")
		w.WriteSeparator()
		w.Indent()
	}
	args, err := litsql.ExpressSlice(w, d, start, c.Set, nil, expr.CommaWriterSeparator, nil)
	if c.Starter {
		w.Dedent()
	}
	return args, err
}

var _ litsql.QueryClauseMerge = (*Set)(nil)

func (c *Set) ClauseID() string {
	return "30279c84-ab70-4b63-aab5-03a7ba8fc111"
}

func (c *Set) ClauseOrder() int {
	return clause.OrderSet
}

func (c *Set) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*Set)
	if !ok {
		panic("invalid merge")
	}
	if o.Starter {
		c.Starter = true
	}
	c.Set = append(c.Set, o.Set...)
}
