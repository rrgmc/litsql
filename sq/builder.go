package sq

import (
	"cmp"
	"slices"

	"github.com/rrgmc/litsql"
)

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

func (s *Builder) Add(q litsql.QueryClause) {
	if w, ok := q.(litsql.QueryClauseWrapper); ok {
		q = w.WrappedQueryClause()
	}
	cid := q.ClauseID()
	if e, ok := s.mlist[cid]; ok {
		if len(e) == 0 {
			panic("should never be 0")
		}
		if em, ok := e[0].(litsql.QueryClauseMerge); ok {
			em.ClauseMerge(q)
		} else if em, ok := e[0].(litsql.QueryClauseMultiple); ok {
			s.mlist[cid] = append(s.mlist[cid], em)
		} else {
			s.mlist[cid][0] = q
		}
		return
	}
	s.mlist[cid] = []litsql.QueryClause{q}
}

func (s *Builder) ClauseList() []litsql.QueryClause {
	var exprs []litsql.QueryClause
	for _, q := range s.mlist {
		exprs = append(exprs, q...)
	}
	slices.SortFunc(exprs, func(a, b litsql.QueryClause) int {
		return cmp.Compare(a.ClauseOrder(), b.ClauseOrder())
	})
	return exprs
}
