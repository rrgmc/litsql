package expr

import (
	"github.com/rrgmc/litsql"
)

func J(exprs ...litsql.Expression) litsql.Expression {
	return join{exprs: exprs, sep: ""}
}

func JS(sep string, exprs ...litsql.Expression) litsql.Expression {
	return join{exprs: exprs, sep: sep}
}

func Or(expr ...string) litsql.Expression {
	return OrE(SL(expr)...)
}

func OrE(expr ...litsql.Expression) litsql.Expression {
	return join{exprs: expr, sep: " OR "}
}

func And(expr ...string) litsql.Expression {
	return AndE(SL(expr)...)
}

func AndE(expr ...litsql.Expression) litsql.Expression {
	return join{exprs: expr, sep: " OR "}
}

type join struct {
	exprs []litsql.Expression
	sep   string
}

func (s join) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	return litsql.ExpressSlice(w, d, start, s.exprs, nil, Raw(s.sep), nil)
}
