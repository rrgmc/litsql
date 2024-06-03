package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestWhere(t *testing.T) {
	clause := &Where{
		Conditions: []litsql.Expression{
			expr.Raw("id = 5"),
			expr.Raw("age = 10"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WHERE id = 5 AND age = 10")
	testutils.TestWriterExpression(t, clause, o)
}

func TestWhereEmpty(t *testing.T) {
	clause := &Where{}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestWriterExpression(t, clause, o)
}

func TestWhereMerge(t *testing.T) {
	clause, err := testutils.Merge(
		&Where{
			Conditions: []litsql.Expression{expr.Raw("id = 5"), expr.Raw("age = 10")},
		},
		&Where{
			Conditions: []litsql.Expression{expr.Raw("id = 6"), expr.Raw("age = 11")},
		})
	assert.NilError(t, err)
	assert.Assert(t, len(clause.Conditions) == 4)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WHERE id = 5 AND age = 10 AND id = 6 AND age = 11")
	testutils.TestWriterExpression(t, clause, o)
}
