package expr

import "github.com/rrgmc/litsql"

// Func calls the function to return the expression to output.
func Func(f func() (litsql.Expression, error)) litsql.Expression {
	return exprFunc(f)
}

type exprFunc func() (litsql.Expression, error)

func (e exprFunc) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) (args []any, err error) {
	ex, err := e()
	if err != nil {
		return nil, err
	}
	return ex.WriteSQL(w, d, start)
}
