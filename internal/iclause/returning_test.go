package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestReturning(t *testing.T) {
	clause := &Returning{
		Expressions: []litsql.Expression{
			expr.Raw("id"),
			expr.Raw("name"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("RETURNING id, name")
	testutils.TestWriterExpression(t, clause, o)
}

func TestReturningEmpty(t *testing.T) {
	clause := &Returning{}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("RETURNING *")
	testutils.TestWriterExpression(t, clause, o)
}

func TestReturningMerge(t *testing.T) {
	clause, err := testutils.Merge(
		&Returning{
			Expressions: []litsql.Expression{expr.Raw("id"), expr.Raw("id2")},
		},
		&Returning{
			Expressions: []litsql.Expression{expr.Raw("id3"), expr.Raw("id4")},
		})
	assert.NilError(t, err)
	assert.Assert(t, len(clause.Expressions) == 4)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("RETURNING id, id2, id3, id4")
	testutils.TestWriterExpression(t, clause, o)
}
