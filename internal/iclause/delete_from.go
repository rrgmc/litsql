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

func (c *DeleteFrom) ClauseID() string {
	return "da91e158-2a88-4220-9a67-910db2a98017"
}

func (c *DeleteFrom) ClauseOrder() int {
	return clause.OrderDeleteFrom
}
