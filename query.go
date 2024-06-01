package litsql

type Query interface {
	Expression
	WriteQuery(w Writer, start int) (args []any, err error)
}

type QueryBuilder interface {
	Dialect() Dialect
	Add(q QueryClause)
}

type BuildQuery interface {
	Build(writerOptions ...WriterOption) (string, []any, error)
}

type QueryClause interface {
	Expression
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

type QueryClauseWrapper interface {
	WrappedQueryClause() QueryClause
}
