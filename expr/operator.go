package expr

import (
	"github.com/rrgmc/litsql"
)

// Join combines a list of [litsql.Expression] in a single expression, without a separator.
func Join(exprs ...litsql.Expression) litsql.Expression {
	return join{expressions: exprs, sep: ""}
}

// JoinSep combines a list of [litsql.Expression] in a single expression, using the passed separator.
func JoinSep(sep string, exprs ...litsql.Expression) litsql.Expression {
	return join{expressions: exprs, sep: sep}
}

// Or outputs the list of expressions separated by " OR ".
func Or(expr ...string) litsql.Expression {
	return OrExpr(StringList(expr)...)
}

// OrExpr outputs the list of expressions separated by " OR ".
func OrExpr(expr ...litsql.Expression) litsql.Expression {
	return join{expressions: expr, sep: " OR "}
}

// And outputs the list of expressions separated by " AND ".
func And(expr ...string) litsql.Expression {
	return AndExpr(StringList(expr)...)
}

// AndExpr outputs the list of expressions separated by " AND ".
func AndExpr(expr ...litsql.Expression) litsql.Expression {
	return join{expressions: expr, sep: " AND "}
}

type join struct {
	expressions []litsql.Expression
	sep         string
}

func (s join) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	return litsql.ExpressSlice(w, d, start, s.expressions, nil, Raw(s.sep), nil)
}
