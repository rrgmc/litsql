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

func (s *Combine) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if s.Strategy == "" {
		return nil, ErrNoCombinationStrategy
	}
	w.AddSeparator(true)
	w.Write(s.Strategy)
	if s.All {
		w.Write(" ALL ")
	} else {
		w.Write(" ")
	}
	return litsql.Express(w, d, start, s.Query)
}

var _ litsql.QueryClauseMultiple = (*Combine)(nil)

func (c *Combine) ClauseID() string {
	return "143a4da4-4963-4cc9-ade3-b7437f93b660"
}

func (s *Combine) ClauseOrder() int {
	return clause.OrderUnion
}

func (s *Combine) ClauseMultiple() {

}
