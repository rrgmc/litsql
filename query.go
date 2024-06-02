package litsql

// Query is the base interface for queries.
type Query interface {
	Expression
	WriteQuery(w Writer, start int) (args []any, err error)
}

// QueryBuilder is the base interface for queries built by lists of clauses.
type QueryBuilder interface {
	Dialect() Dialect
	AddQueryClause(q QueryClause)
}

// QueryClause is a query clause.
type QueryClause interface {
	Expression
	ClauseID() string
	ClauseOrder() int
}

// QueryClauseMerge can be implemented by QueryClause when its data can be merged.
type QueryClauseMerge interface {
	QueryClause
	ClauseMerge(other QueryClause)
}

// QueryClauseMultiple can be implemented by QueryClause to signal multiple instances can be added.
type QueryClauseMultiple interface {
	QueryClause
	ClauseMultiple()
}

// helpers

// QueryFunc is a functional implementation of Query.
type QueryFunc struct {
	D Dialect
	E Expression
	F func(w Writer, start int) (args []any, err error) // if nil, WriteSQL will be called directly.
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
