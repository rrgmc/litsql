package iclause

import (
	"errors"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq/clause"
)

var ErrNoCombinationStrategy = errors.New("Combination strategy must be set")

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

func (s *Combine) ClauseOrder() int {
	return clause.OrderUnion
}

func (s *Combine) ClauseMultiple() {

}
