package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/sq/clause"
)

type WindowDef struct {
	From        string // an existing window name
	OrderBy     []litsql.Expression
	PartitionBy []litsql.Expression
	Frame
}

func (wd *WindowDef) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, d, start)

	prefixCond := false

	if wd.From != "" {
		w.Write(wd.From)
		prefixCond = true
	}

	b.ExpressSlice(wd.PartitionBy, expr.PrefixIf(prefixCond, expr.Space, expr.Raw("PARTITION BY ")), expr.CommaSpace, nil)
	prefixCond = prefixCond || len(wd.PartitionBy) > 0

	b.ExpressSlice(wd.OrderBy, expr.PrefixIf(prefixCond, expr.Space, expr.Raw("ORDER BY ")), expr.CommaSpace, nil)
	prefixCond = prefixCond || len(wd.OrderBy) > 0

	b.ExpressIf(&wd.Frame, wd.Frame.Defined, expr.Space, nil)

	return b.Result()
}

func (wd *WindowDef) SetFrom(from string) {
	wd.From = from
}

func (wd *WindowDef) AddPartitionBy(condition ...litsql.Expression) {
	wd.PartitionBy = append(wd.PartitionBy, condition...)
}

func (wd *WindowDef) AddOrderBy(order ...litsql.Expression) {
	wd.OrderBy = append(wd.OrderBy, order...)
}

type NamedWindow struct {
	Name       string
	Definition WindowDef
}

func (n *NamedWindow) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if n.Name == "" {
		return nil, internal.NewClauseError("'WINDOW' name is required")
	}

	w.Write(n.Name)
	w.Write(" AS (")
	args, err := litsql.Express(w, d, start, &n.Definition)
	w.Write(internal.ClosePar)
	return args, err
}

type Windows struct {
	Windows []*NamedWindow
}

func (c *Windows) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if len(c.Windows) == 0 {
		return nil, nil
	}

	b := litsql.NewExpressBuilder(w, d, start)

	w.AddSeparator(true)
	w.Write("WINDOW ")

	if len(c.Windows) > 1 {
		w.WriteNewLine()
		w.Indent()
	}
	b.ExpressSlice(expr.CastSlice(c.Windows), nil, expr.CommaWriterSeparator, nil)
	if len(c.Windows) > 1 {
		w.Dedent()
	}

	return b.Result()
}

var _ litsql.QueryClauseMerge = (*Windows)(nil)

func (c *Windows) ClauseID() string {
	return "d8c54a38-29e2-4205-a991-e24c8238ae00"
}

func (c *Windows) ClauseOrder() int {
	return clause.OrderWindow
}

func (c *Windows) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*Windows)
	if !ok {
		panic("invalid merge")
	}
	c.Windows = append(c.Windows, o.Windows...)
}
