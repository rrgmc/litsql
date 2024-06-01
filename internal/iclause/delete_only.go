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

func (c *DeleteOnly) ClauseOrder() int {
	return clause.OrderDeleteOnly
}
