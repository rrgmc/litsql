package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestColumns(t *testing.T) {
	c := &Columns{
		Columns: []litsql.Expression{
			expr.Raw("id"),
			expr.Raw("name"),
		},
	}

	o := testutils.NewTestBuffer()
	o.Write("id, name")
	testutils.CheckExpression(t, c, o)
}
