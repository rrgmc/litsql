package iclause

import (
	"errors"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type Values struct {
	// Query takes the highest priority
	// If present, will attempt to insert from this query
	Query litsql.Query

	// for multiple inserts
	// each sub-slice is one set of values
	Vals []Value
}

func (c *Values) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	// If a query is present, use it
	if c.Query != nil {
		if len(c.Vals) > 0 {
			return nil, errors.New("cannot set Query and Values at the same time in Values clause")
		}
		w.WriteSeparator()
		return c.Query.WriteQuery(w, start)
	}

	w.AddSeparator(true)

	// If values are present, use them
	if len(c.Vals) > 0 {
		w.Write("VALUES")
		if len(c.Vals) > 1 {
			w.WriteSeparator()
			w.Indent()
		} else {
			w.Write(" ")
		}
		args, err := litsql.ExpressSlice(w, d, start, expr.CastSlice(c.Vals), nil, expr.CommaWriterSeparator, nil)
		if len(c.Vals) > 1 {
			w.Dedent()
		}
		return args, err
	}

	// If no value was present, use default value
	w.Write("DEFAULT VALUES")
	return nil, nil
}

var _ litsql.QueryClauseMerge = (*Values)(nil)

func (c *Values) ClauseID() string {
	return "40ca9400-2878-49c8-8c27-9aef6b75392c"
}

func (c *Values) ClauseOrder() int {
	return clause.OrderValues
}

func (c *Values) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*Values)
	if !ok {
		panic("invalid merge")
	}
	if o.Query != nil {
		c.Query = o.Query
	}
	c.Vals = append(c.Vals, o.Vals...)
}

type Value []litsql.Expression

func (v Value) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	return litsql.ExpressSlice(w, d, start, v, expr.OpenPar, expr.CommaSpace, expr.ClosePar)
}
