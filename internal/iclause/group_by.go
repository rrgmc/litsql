package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
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

func (c *GroupBy) ClauseID() string {
	return "fc5be717-8bed-4350-bcfe-4d6a0b362506"
}

func (c *GroupBy) ClauseOrder() int {
	return clause.OrderGroupBy
}

func (c *GroupBy) ClauseMerge(other litsql.QueryClause) error {
	o, ok := other.(*GroupBy)
	if !ok {
		return internal.NewClauseErrorInvalidMerge("GroupBy")
	}
	if c.Distinct != o.Distinct || o.With != "" {
		return internal.NewClauseErrorInvalidMergeHasChanges("GroupBy")
	}
	c.Groups = append(c.Groups, o.Groups...)
	return nil
}

func (g *GroupBy) SetGroupWith(with string) {
	g.With = with
}

func (g *GroupBy) SetGroupByDistinct(distinct bool) {
	g.Distinct = distinct
}
