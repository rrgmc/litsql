package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq/clause"
)

type DeleteFrom struct {
	Table string
}

func (c *DeleteFrom) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	w.Write(c.Table)
	return nil, nil
}

var _ litsql.QueryClause = (*DeleteFrom)(nil)

func (c *DeleteFrom) ClauseOrder() int {
	return clause.OrderDeleteFrom
}
