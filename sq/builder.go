package sq

import (
	"cmp"
	"errors"
	"fmt"
	"slices"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

// Builder implements [litsql.QueryBuilder].
type Builder struct {
	d      litsql.Dialect
	mlist  map[string][]litsql.QueryClause
	addErr error
}

func NewQueryBuilder(d litsql.Dialect) *Builder {
	return &Builder{
		d:     d,
		mlist: make(map[string][]litsql.QueryClause),
	}
}

func (s *Builder) Dialect() litsql.Dialect {
	return s.d
}

func (s *Builder) AddQueryClause(q litsql.QueryClause) {
	cid := q.ClauseID()
	s.mlist[cid] = append(s.mlist[cid], q)
}

func (s *Builder) QueryClauseList() ([]litsql.QueryClause, error) {
	if s.addErr != nil {
		return nil, s.addErr
	}

	var exprs []litsql.QueryClause
	for _, e := range s.mlist {
		if len(e) == 0 {
			return nil, errors.New("invalid condition: clause list should never be 0")
		}
		if _, ok := e[0].(litsql.QueryClauseMerge); ok {
			// clause can be merged, merge on the first one.
			em, err := internal.MergeClauses(e...)
			if err != nil {
				return nil, fmt.Errorf("error merging clauses: %w", err)
			}
			exprs = append(exprs, em)
		} else if _, ok := e[0].(litsql.QueryClauseMultiple); ok {
			// clause can have multiple instances.
			exprs = append(exprs, e...)
		} else {
			// clause cannot have multiple instances
			if len(e) > 1 {
				return nil, internal.NewClauseErrorInvalidMergeCannotHaveMultiple(fmt.Sprintf("%T", e[0]))
			}
			exprs = append(exprs, e[0])
		}
	}
	slices.SortFunc(exprs, func(a, b litsql.QueryClause) int {
		return cmp.Compare(a.ClauseOrder(), b.ClauseOrder())
	})
	return exprs, nil
}
