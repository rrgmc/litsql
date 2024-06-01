package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
)

func TestColumns(t *testing.T) {
	c := &Columns{
		Columns: []litsql.Expression{
			expr.Raw("id"),
			expr.Raw("name"),
		},
	}

	var o testBuffer
	o.Write("id, name")
	checkExpression(t, c, o)
}
