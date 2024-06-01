package expr

import (
	"github.com/rrgmc/litsql"
)

func Paren(expr ...string) litsql.Expression {
	return ParenE(SL(expr)...)
}

func ParenE(expr ...litsql.Expression) litsql.Expression {
	return group(expr)
}

type group []litsql.Expression

func (g group) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if len(g) == 0 {
		return litsql.ExpressIf(w, d, start, null, true, OpenPar, ClosePar)
	}
	return litsql.ExpressSlice(w, d, start, g, OpenPar, CommaSpace, ClosePar)
}
