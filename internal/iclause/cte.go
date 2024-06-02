package iclause

import (
	"fmt"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
)

type CTE struct {
	Query        litsql.Query // SQL standard says only select, postgres allows insert/update/delete
	Name         string
	Columns      []litsql.Expression
	Materialized *bool
	Search       CTESearch
	Cycle        CTECycle
}

func (c *CTE) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if c.Name == "" {
		return nil, internal.NewClauseError("'WITH' CTE name is required")
	}

	b := litsql.NewExpressBuilder(w, d, start)

	w.Write(c.Name)
	b.ExpressSlice(c.Columns, expr.OpenPar, expr.CommaSpace, expr.ClosePar)
	w.Write(" AS ")

	switch {
	case c.Materialized == nil:
		// do nothing
		break
	case *c.Materialized:
		w.Write("MATERIALIZED ")
	case !*c.Materialized:
		w.Write("NOT MATERIALIZED ")
	}

	b.Express(c.Query)

	b.ExpressIf(&c.Search, len(c.Search.Columns) > 0, expr.WriterSeparator, nil)
	b.ExpressIf(&c.Cycle, len(c.Cycle.Columns) > 0, expr.WriterSeparator, nil)

	return b.Result()
}

func (c *CTE) SetQuery(q litsql.Query) {
	c.Query = q
}

func (c *CTE) SetNotMaterialized() {
	c.Materialized = new(bool)
	*c.Materialized = false
}

func (c *CTE) SetMaterialized() {
	c.Materialized = new(bool)
	*c.Materialized = true
}

const (
	SearchBreadth = "BREADTH"
	SearchDepth   = "DEPTH"
)

type CTESearch struct {
	Order   string
	Columns []string
	Set     string
}

func (c *CTESearch) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	// [ SEARCH { BREADTH | DEPTH } FIRST BY column_name [, ...] SET search_seq_col_name ]
	w.Write(fmt.Sprintf("SEARCH %s FIRST BY ", c.Order))

	args, err := litsql.ExpressSlice(w, d, start, expr.SL(c.Columns), nil, expr.CommaSpace, nil)
	if err != nil {
		return nil, err
	}

	w.Write(fmt.Sprintf(" SET %s", c.Set))

	return args, nil
}

type CTECycle struct {
	Columns    []string
	Set        string
	Using      string
	SetVal     litsql.Expression
	DefaultVal litsql.Expression
}

func (c *CTECycle) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, d, start)

	// [ CYCLE column_name [, ...] SET cycle_mark_col_name [ TO cycle_mark_value DEFAULT cycle_mark_default ] USING cycle_path_col_name ]
	w.Write("CYCLE ")
	b.ExpressSlice(expr.SL(c.Columns), nil, expr.CommaSpace, nil)

	w.Write(fmt.Sprintf(" SET %s", c.Set))
	b.ExpressIf(c.SetVal, c.SetVal != nil, expr.Raw(" TO "), nil)
	b.ExpressIf(c.DefaultVal, c.DefaultVal != nil, expr.Raw(" DEFAULT "), nil)

	w.Write(" USING ")
	w.Write(c.Using)

	return b.Result()
}
