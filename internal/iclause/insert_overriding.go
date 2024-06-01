package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq/clause"
)

type InsertOverriding struct {
	Overriding string
}

func (c *InsertOverriding) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	return litsql.ExpressIf(w, d, start, expr.Raw(c.Overriding),
		c.Overriding != "", expr.J(expr.WriterAddSeparator(true), expr.Raw("OVERRIDING ")), expr.Raw(" VALUE"))
}

var _ litsql.QueryClause = (*InsertOverriding)(nil)

func (c *InsertOverriding) ClauseOrder() int {
	return clause.OrderInsertOverriding
}
