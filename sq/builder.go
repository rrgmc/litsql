package sq

import (
	"cmp"
	"reflect"
	"slices"

	"github.com/rrgmc/litsql"
)

type Builder struct {
	d     litsql.Dialect
	mlist map[reflect.Type][]litsql.QueryClause
}

func NewQueryBuilder(d litsql.Dialect) *Builder {
	return &Builder{
		d:     d,
		mlist: make(map[reflect.Type][]litsql.QueryClause),
	}
}

func (s *Builder) Dialect() litsql.Dialect {
	return s.d
}

func (s *Builder) Add(q litsql.QueryClause) {
	if w, ok := q.(litsql.QueryClauseWrapper); ok {
		q = w.WrappedQueryClause()
	}
	tp := reflect.TypeOf(q)
	if e, ok := s.mlist[tp]; ok {
		if len(e) == 0 {
			panic("should never be 0")
		}
		if em, ok := e[0].(litsql.QueryClauseMerge); ok {
			em.ClauseMerge(q)
		} else if em, ok := e[0].(litsql.QueryClauseMultiple); ok {
			s.mlist[tp] = append(s.mlist[tp], em)
		} else {
			s.mlist[tp][0] = q
		}
		return
	}
	s.mlist[tp] = []litsql.QueryClause{q}
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
