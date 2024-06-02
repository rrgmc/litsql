package sq

import (
	"cmp"
	"slices"

	"github.com/rrgmc/litsql"
)

// Builder implements [litsql.QueryBuilder].
type Builder struct {
	d     litsql.Dialect
	mlist map[string][]litsql.QueryClause
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
	if e, ok := s.mlist[cid]; ok {
		if len(e) == 0 {
			panic("should never be 0")
		}
		if em, ok := e[0].(litsql.QueryClauseMerge); ok {
			// clause can be merged, merge on the first one.
			em.ClauseMerge(q)
		} else if em, ok := e[0].(litsql.QueryClauseMultiple); ok {
			// clause can have multiple instances.
			s.mlist[cid] = append(s.mlist[cid], em)
		} else {
			// clause can have only a single instance, store only the last one.
			s.mlist[cid][0] = q
		}
		return
	}
	s.mlist[cid] = []litsql.QueryClause{q}
}

func (s *Builder) QueryClauseList() []litsql.QueryClause {
	var exprs []litsql.QueryClause
	for _, q := range s.mlist {
		exprs = append(exprs, q...)
	}
	slices.SortFunc(exprs, func(a, b litsql.QueryClause) int {
		return cmp.Compare(a.ClauseOrder(), b.ClauseOrder())
	})
	return exprs
}
