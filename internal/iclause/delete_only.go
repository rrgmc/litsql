package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq/clause"
)

type DeleteOnly struct {
	Only bool
}

func (c *DeleteOnly) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if c.Only {
		w.Write("ONLY")
	}
	return nil, nil
}

var _ litsql.QueryClause = (*DeleteOnly)(nil)

func (c *DeleteOnly) ClauseID() string {
	return "f1c07cf2-1b54-4f83-91a4-b5420514b373"
}

func (c *DeleteOnly) ClauseOrder() int {
	return clause.OrderDeleteOnly
}
