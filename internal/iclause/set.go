package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/sq/clause"
)

type Set struct {
	Set     []litsql.Expression
	Starter bool
}

func (c *Set) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if c.Starter {
		if len(c.Set) == 0 {
			return nil, internal.NewClauseError("'SET' fields are required")
		}
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

func (c *Set) ClauseMerge(other litsql.QueryClause) error {
	o, ok := other.(*Set)
	if !ok {
		return internal.NewClauseErrorInvalidMerge("Set")
	}
	if c.Starter != o.Starter {
		return internal.NewClauseErrorInvalidMergeHasChanges("Set")
	}
	c.Set = append(c.Set, o.Set...)
	return nil
}
