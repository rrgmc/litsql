package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq/clause"
)

type RawQuery struct {
	Query litsql.Expression
	Args  []any
}

func (c *RawQuery) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	args, err := litsql.Express(w, d, start, c.Query)
	if err != nil {
		return nil, err
	}
	return append(args, c.Args...), nil
}

var _ litsql.QueryClause = (*RawQuery)(nil)

func (c *RawQuery) ClauseID() string {
	return "79bcdb4d-c414-4989-954e-99f7c5ed2bfa"
}

func (c *RawQuery) ClauseOrder() int {
	return clause.OrderRawQuery
}
