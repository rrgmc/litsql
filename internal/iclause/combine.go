package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/sq/clause"
)

var ErrNoCombinationStrategy = internal.NewClauseError("Combination strategy must be set")

const (
	Union     = "UNION"
	Intersect = "INTERSECT"
	Except    = "EXCEPT"
)

type Combine struct {
	Strategy string
	Query    litsql.Query
	All      bool
}

func (c *Combine) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if c.Strategy == "" {
		return nil, ErrNoCombinationStrategy
	}
	w.AddSeparator(true)
	w.Write(c.Strategy)
	if c.All {
		w.Write(" ALL ")
	} else {
		w.Write(" ")
	}
	return litsql.Express(w, d, start, c.Query)
}

var _ litsql.QueryClauseMultiple = (*Combine)(nil)

func (c *Combine) ClauseID() string {
	return "143a4da4-4963-4cc9-ade3-b7437f93b660"
}

func (c *Combine) ClauseOrder() int {
	return clause.OrderUnion
}

func (c *Combine) ClauseMultiple() {

}
