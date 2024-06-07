package isq

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/sq"
)

type DefaultQuery[T any] struct {
	builder    *sq.QueryBuilder
	startOrder int
	startStr   string
}

func NewDefaultQuery[T any](d litsql.Dialect, startOrder int, startStr string) *DefaultQuery[T] {
	return &DefaultQuery[T]{
		builder:    sq.NewQueryBuilder(d),
		startOrder: startOrder,
		startStr:   startStr,
	}
}

func (s *DefaultQuery[T]) Dialect() litsql.Dialect {
	return s.builder.Dialect()
}

func (s *DefaultQuery[T]) AddQueryClause(q litsql.QueryClause) {
	s.builder.AddQueryClause(q)
}

func (s *DefaultQuery[T]) WriteSQL(w litsql.Writer, _ litsql.Dialect, start int) ([]any, error) {
	w.Write(internal.OpenPar)
	w.WriteNewLine()
	w.Indent()
	args, err := s.WriteQuery(w, start)
	w.Dedent()
	w.WriteNewLine()
	w.Write(internal.ClosePar)
	return args, err
}

func (s *DefaultQuery[T]) WriteQuery(w litsql.Writer, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, s.Dialect(), start)

	w.StartQuery()
	wroteStart := false

	clauses, err := s.builder.QueryClauseList()
	if err != nil {
		return nil, err
	}
	for i, e := range clauses {
		if !wroteStart && e.ClauseOrder() >= s.startOrder {
			w.AddSeparator(true)
			w.Write(s.startStr)
			wroteStart = true
		} else if i > 0 {
			w.AddSeparator(false)
		}
		b.Express(e)
	}

	return b.Result()
}

func (s *DefaultQuery[T]) Apply(mods ...sq.QueryMod[T]) {
	for _, m := range mods {
		m.Apply(s)
	}
}

func (s *DefaultQuery[T]) Build(options ...sq.BuildOption) (string, []any, error) {
	return sq.Build(s, options...)
}
