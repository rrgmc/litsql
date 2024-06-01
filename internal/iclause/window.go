package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/sq/clause"
)

type WindowDef struct {
	From        string // an existing window name
	orderBy     []litsql.Expression
	partitionBy []litsql.Expression
	Frame
}

func (wi *WindowDef) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, d, start)

	if wi.From != "" {
		w.Write(wi.From)
		w.Write(" ")
	}

	b.ExpressSlice(wi.partitionBy, expr.Raw("PARTITION BY "), expr.CommaSpace, expr.Space)

	b.ExpressSlice(wi.orderBy, expr.Raw("ORDER BY "), expr.CommaSpace, nil)

	b.ExpressIf(&wi.Frame, wi.Frame.Defined, expr.Space, nil)

	return b.Result()
}

func (wi *WindowDef) SetFrom(from string) {
	wi.From = from
}

func (wi *WindowDef) AddPartitionBy(condition ...litsql.Expression) {
	wi.partitionBy = append(wi.partitionBy, condition...)
}

func (wi *WindowDef) AddOrderBy(order ...litsql.Expression) {
	wi.orderBy = append(wi.orderBy, order...)
}

type NamedWindow struct {
	Name       string
	Definition WindowDef
}

func (n NamedWindow) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	w.Write(n.Name)
	w.Write(" AS (")
	args, err := litsql.Express(w, d, start, &n.Definition)
	w.Write(internal.ClosePar)
	return args, err
}

type Windows struct {
	Windows []NamedWindow
}

func (wi *Windows) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, d, start)

	w.AddSeparator(true)
	w.Write("WINDOW ")

	if len(wi.Windows) > 1 {
		w.WriteNewLine()
		w.Indent()
	}
	b.ExpressSlice(expr.CastSlice(wi.Windows), nil, expr.CommaWriterSeparator, nil)
	if len(wi.Windows) > 1 {
		w.Dedent()
	}

	return b.Result()
}

var _ litsql.QueryClauseMerge = (*Windows)(nil)

func (c *Windows) ClauseID() string {
	return "d8c54a38-29e2-4205-a991-e24c8238ae00"
}

func (wi *Windows) ClauseOrder() int {
	return clause.OrderWindow
}

func (wi *Windows) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*Windows)
	if !ok {
		panic("invalid merge")
	}
	wi.Windows = append(wi.Windows, o.Windows...)
}
