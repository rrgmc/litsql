package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestColumns(t *testing.T) {
	clause := &Columns{
		Columns: []litsql.Expression{
			expr.Raw("id"),
			expr.Raw("name"),
		},
	}

	o := testutils.NewTestBuffer()
	o.Write("id, name")
	testutils.TestWriterExpression(t, clause, o)
}

func TestColumnsEmpty(t *testing.T) {
	clause := &Columns{}

	o := testutils.NewTestBuffer()
	o.Write("*")
	testutils.TestWriterExpression(t, clause, o)
}

func TestColumnsMerge(t *testing.T) {
	clause, err := testutils.Merge(
		&Columns{
			Columns: []litsql.Expression{expr.Raw("id"), expr.Raw("id2")},
		},
		&Columns{
			Columns: []litsql.Expression{expr.Raw("id3"), expr.Raw("id4")},
		})
	assert.NilError(t, err)
	assert.Assert(t, len(clause.Columns) == 4)

	o := testutils.NewTestBuffer()
	o.Write("id, id2, id3, id4")
	testutils.TestWriterExpression(t, clause, o)
}
