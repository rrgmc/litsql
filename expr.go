package litsql

// Expression is the base expression interface.
type Expression interface {
	WriteSQL(w Writer, d Dialect, start int) (args []any, err error)
}

// ExpressionFunc is the functional implementation of Expression.
type ExpressionFunc func(w Writer, d Dialect, start int) ([]any, error)

func (e ExpressionFunc) WriteSQL(w Writer, d Dialect, start int) ([]any, error) {
	return e(w, d, start)
}
