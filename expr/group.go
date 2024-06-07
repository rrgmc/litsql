package expr

import (
	"github.com/rrgmc/litsql"
)

// Paren returns an expression what outputs the list of expressions comma-separated with parenthesis around.
func Paren(expr ...string) litsql.Expression {
	return ParenExpr(StringList(expr)...)
}

// ParenExpr returns an expression what outputs the list of expressions comma-separated with parenthesis around.
func ParenExpr(expr ...litsql.Expression) litsql.Expression {
	return group(expr)
}

type group []litsql.Expression

func (g group) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if len(g) == 0 {
		return litsql.ExpressIf(w, d, start, null, true, OpenPar, ClosePar)
	}
	return litsql.ExpressSlice(w, d, start, g, OpenPar, CommaSpace, ClosePar)
}
