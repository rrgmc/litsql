package litsql

type Query interface {
	Expression
	WriteQuery(w Writer, start int) (args []any, err error)
}

type QueryBuilder interface {
	Dialect() Dialect
	Add(q QueryClause)
}

type QueryClause interface {
	Expression
	ClauseID() string
	ClauseOrder() int
}

type QueryClauseMerge interface {
	QueryClause
	ClauseMerge(other QueryClause)
}

type QueryClauseMultiple interface {
	QueryClause
	ClauseMultiple()
}

// helpers

type QueryFunc struct {
	D Dialect
	E Expression
	F func(w Writer, start int) (args []any, err error)
}

func (q QueryFunc) WriteQuery(w Writer, start int) (args []any, err error) {
	if q.F == nil {
		return q.WriteSQL(w, q.D, start)
	}
	return q.F(w, start)
}

func (q QueryFunc) WriteSQL(w Writer, d Dialect, start int) (args []any, err error) {
	return q.E.WriteSQL(w, d, start)
}
