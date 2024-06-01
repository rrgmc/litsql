package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq/clause"
)

type UpdateOnly struct {
	Only bool
}

func (c *UpdateOnly) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if c.Only {
		w.Write("ONLY")
	}
	return nil, nil
}

var _ litsql.QueryClause = (*UpdateOnly)(nil)

func (c *UpdateOnly) ClauseID() string {
	return "d62f9d23-74a7-40cb-b7bd-4eefd76cc630"
}

func (c *UpdateOnly) ClauseOrder() int {
	return clause.OrderUpdateOnly
}
