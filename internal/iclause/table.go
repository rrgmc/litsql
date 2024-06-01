package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq/clause"
)

type Table struct {
	Expression litsql.Expression
	Alias      string
	Columns    []string
}

func (t *Table) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, d, start)

	b.Express(t.Expression)

	if t.Alias != "" {
		w.Write(" AS ")
		w.Write(t.Alias)
	}

	if len(t.Columns) > 0 {
		w.Write(" (")
		for k, cAlias := range t.Columns {
			if k != 0 {
				w.Write(", ")
			}
			w.Write(cAlias)
		}
		w.Write(")")
	}

	return b.Result()
}

var _ litsql.QueryClause = (*Table)(nil)

func (c *Table) ClauseOrder() int {
	return clause.OrderTable
}
