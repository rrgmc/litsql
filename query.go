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
	ClauseMerge(other QueryClause) error
}

// QueryClauseMultiple can be implemented by QueryClause to signal multiple instances can be added.
type QueryClauseMultiple interface {
	QueryClause
	ClauseMultiple()
}

// helpers

// QueryFunc is a functional implementation of Query.
// If f is nil, WriteSQL will be called directly by WriteQuery.
func QueryFunc(dialect Dialect, expression Expression, f func(w Writer, start int) (args []any, err error)) Query {
	return &queryFunc{dialect, expression, f}
}

// queryFunc is a functional implementation of Query.
type queryFunc struct {
	dialect    Dialect
	expression Expression
	fn         func(w Writer, start int) (args []any, err error) // if nil, WriteSQL will be called directly by WriteQuery.
}

func (q queryFunc) WriteQuery(w Writer, start int) (args []any, err error) {
	if q.fn == nil {
		return q.WriteSQL(w, q.dialect, start)
	}
	return q.fn(w, start)
}

func (q queryFunc) WriteSQL(w Writer, d Dialect, start int) (args []any, err error) {
	return q.expression.WriteSQL(w, d, start)
}
