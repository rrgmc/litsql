package testutils

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
)

func TestClausePrefix(e litsql.Expression) litsql.Expression {
	return expr.J(expr.Raw("@"), e)
}
