package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestWith(t *testing.T) {
	clause := &With{
		CTEs: []*CTE{
			{
				Name: "testCTE",
				Columns: []litsql.Expression{
					expr.Raw("id"),
					expr.Raw("name"),
				},
				Query: litsql.QueryFunc(testutils.NewTestDialect(), expr.Raw("test_query"), nil),
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WITH testCTE(id, name) AS test_query")
	testutils.TestWriterExpression(t, clause, o)
}

func TestWithEmpty(t *testing.T) {
	clause := &With{}
	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestWriterExpression(t, clause, o)
}

func TestWithEmptyCTE(t *testing.T) {
	clause := &With{
		CTEs: []*CTE{
			{},
		},
	}

	testutils.TestWriterExpressionErrorIs(t, clause, litsql.ErrClause)
}

func TestWithMerge(t *testing.T) {
	clause, err := internal.MergeClauses(
		&With{
			CTEs: []*CTE{
				{
					Name: "testCTE",
					Columns: []litsql.Expression{
						expr.Raw("id"),
						expr.Raw("name"),
					},
					Query: litsql.QueryFunc(testutils.NewTestDialect(), expr.Raw("test_query"), nil),
				},
			},
		},
		&With{
			CTEs: []*CTE{
				{
					Name: "testCTE2",
					Columns: []litsql.Expression{
						expr.Raw("id"),
						expr.Raw("name"),
					},
					Query: litsql.QueryFunc(testutils.NewTestDialect(), expr.Raw("test_query2"), nil),
				},
			},
		})
	assert.NilError(t, err)
	assert.Assert(t, len(clause.CTEs) == 2)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WITH testCTE(id, name) AS test_query,")
	o.WriteSeparator()
	o.Write("testCTE2(id, name) AS test_query2")
	testutils.TestWriterExpression(t, clause, o)
}

func TestWithMergeRecursive(t *testing.T) {
	rtrue := true
	rfalse := false

	_, err := internal.MergeClauses(
		&With{
			Recursive: &rtrue,
			CTEs: []*CTE{
				{
					Name: "testCTE",
					Columns: []litsql.Expression{
						expr.Raw("id"),
						expr.Raw("name"),
					},
					Query: litsql.QueryFunc(testutils.NewTestDialect(), expr.Raw("test_query"), nil),
				},
			},
		},
		&With{
			Recursive: &rfalse,
			CTEs: []*CTE{
				{
					Name: "testCTE2",
					Columns: []litsql.Expression{
						expr.Raw("id"),
						expr.Raw("name"),
					},
					Query: litsql.QueryFunc(testutils.NewTestDialect(), expr.Raw("test_query2"), nil),
				},
			},
		})
	assert.ErrorIs(t, err, litsql.ErrClause)
}
