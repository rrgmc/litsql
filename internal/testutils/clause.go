package testutils

import (
	"github.com/rrgmc/litsql"
)

func testClausePrefix(e litsql.Expression) litsql.Expression {
	return litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		w.Write("@")
		return e.WriteSQL(w, d, start)
	})
}
