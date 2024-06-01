package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type GroupBy struct {
	Groups   []litsql.Expression
	Distinct bool
	With     string // ROLLUP | CUBE
}

func (g *GroupBy) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if len(g.Groups) == 0 {
		return nil, nil
	}

	w.AddSeparator(true)
	w.Write("GROUP BY ")
	if g.Distinct {
		w.Write("DISTINCT ")
	}

	args, err := litsql.ExpressSlice(w, d, start, g.Groups, nil, expr.CommaSpace, nil)
	if err != nil {
		return nil, err
	}

	if g.With != "" {
		w.Write(" WITH ")
		w.Write(g.With)
	}

	return args, nil
}

var _ litsql.QueryClauseMerge = (*GroupBy)(nil)

func (c *GroupBy) ClauseOrder() int {
	return clause.OrderGroupBy
}

func (c *GroupBy) ClauseMerge(other litsql.QueryClause) {
	o, ok := other.(*GroupBy)
	if !ok {
		panic("invalid merge")
	}
	c.Groups = append(c.Groups, o.Groups...)
}

func (g *GroupBy) SetGroupWith(with string) {
	g.With = with
}

func (g *GroupBy) SetGroupByDistinct(distinct bool) {
	g.Distinct = distinct
}
