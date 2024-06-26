package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestSet(t *testing.T) {
	clause := &Set{
		Starter: true,
		Set: []litsql.Expression{
			expr.Raw("id = 1"),
			expr.Raw("age = 50"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("SET")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("id = 1,")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("age = 50")
	testutils.TestWriterExpression(t, clause, o)
}

func TestSetEmpty(t *testing.T) {
	clause := &Set{}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestWriterExpression(t, clause, o)
}

func TestSetEmptyStarter(t *testing.T) {
	clause := &Set{
		Starter: true,
	}
	testutils.TestWriterExpressionErrorIs(t, clause, litsql.ErrClause)
}

func TestSetMerge(t *testing.T) {
	clause, err := internal.MergeClauses(
		&Set{
			Starter: true,
			Set:     []litsql.Expression{expr.Raw("id = 5"), expr.Raw("age = 10")},
		},
		&Set{
			Starter: true,
			Set:     []litsql.Expression{expr.Raw("id = 6"), expr.Raw("age = 11")},
		})
	assert.NilError(t, err)
	assert.Assert(t, len(clause.Set) == 4)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("SET")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("id = 5,")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("age = 10,")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("id = 6,")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("age = 11")
	testutils.TestWriterExpression(t, clause, o)
}
