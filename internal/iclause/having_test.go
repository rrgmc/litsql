package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestHaving(t *testing.T) {
	clause := &Having{
		Conditions: []litsql.Expression{
			expr.Raw("id = 5"),
			expr.Raw("age = 10"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("HAVING id = 5 AND age = 10")
	testutils.TestExpression(t, clause, o)
}

func TestHavingEmpty(t *testing.T) {
	clause := &Having{}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestExpression(t, clause, o)
}

func TestHavingMerge(t *testing.T) {
	clause := testutils.Merge(
		&Having{
			Conditions: []litsql.Expression{expr.Raw("id = 5"), expr.Raw("age = 10")},
		},
		&Having{
			Conditions: []litsql.Expression{expr.Raw("id = 6"), expr.Raw("age = 11")},
		})
	assert.Assert(t, len(clause.Conditions) == 4)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("HAVING id = 5 AND age = 10 AND id = 6 AND age = 11")
	testutils.TestExpression(t, clause, o)
}
