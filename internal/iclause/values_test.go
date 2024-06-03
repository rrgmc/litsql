package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestValues(t *testing.T) {
	clause := &Values{
		Vals: []Value{
			{
				expr.Raw("5"),
				expr.Raw("50"),
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("VALUES (5, 50)")
	testutils.TestWriterExpression(t, clause, o)
}

func TestValuesEmpty(t *testing.T) {
	clause := &Values{}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("DEFAULT VALUES")
	testutils.TestWriterExpression(t, clause, o)
}

func TestValuesMultiple(t *testing.T) {
	clause := &Values{
		Vals: []Value{
			{
				expr.Raw("5"),
				expr.Raw("50"),
			},
			{
				expr.Raw("100"),
				expr.Raw("200"),
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("VALUES")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("(5, 50),")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("(100, 200)")
	testutils.TestWriterExpression(t, clause, o)
}

func TestValuesQuery(t *testing.T) {
	clause := &Values{
		Query: litsql.QueryFunc{
			D: testutils.NewTestDialect(),
			E: expr.Raw("test_query"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("test_query")
	testutils.TestWriterExpression(t, clause, o)
}

func TestValuesQueryAndValues(t *testing.T) {
	clause := &Values{
		Query: litsql.QueryFunc{
			D: testutils.NewTestDialect(),
			E: expr.Raw("test_query"),
		},
		Vals: []Value{
			{
				expr.Raw("5"),
				expr.Raw("50"),
			},
		},
	}
	testutils.TestWriterExpressionErrorIs(t, clause, litsql.ErrClause)
}

func TestValuesMerge(t *testing.T) {
	clause, err := internal.MergeClauses(
		&Values{
			Vals: []Value{
				{
					expr.Raw("5"),
					expr.Raw("50"),
				},
			},
		},
		&Values{
			Vals: []Value{
				{
					expr.Raw("100"),
					expr.Raw("200"),
				},
			},
		})
	assert.NilError(t, err)
	assert.Assert(t, len(clause.Vals) == 2)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("VALUES")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("(5, 50),")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("(100, 200)")
	testutils.TestWriterExpression(t, clause, o)
}
