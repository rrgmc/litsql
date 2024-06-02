package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
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
				Query: litsql.QueryFunc{
					D: testutils.NewTestDialect(),
					E: expr.Raw("test_query"),
				},
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WITH testCTE(id, name) AS test_query")
	testutils.TestExpression(t, clause, o)
}

func TestWithEmpty(t *testing.T) {
	clause := &With{}
	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestExpression(t, clause, o)
}

func TestWithEmptyCTE(t *testing.T) {
	clause := &With{
		CTEs: []*CTE{
			{},
		},
	}

	testutils.TestExpressionErrorIs(t, clause, litsql.ErrClause)
}

func TestWithMerge(t *testing.T) {
	clause := testutils.Merge(
		&With{
			CTEs: []*CTE{
				{
					Name: "testCTE",
					Columns: []litsql.Expression{
						expr.Raw("id"),
						expr.Raw("name"),
					},
					Query: litsql.QueryFunc{
						D: testutils.NewTestDialect(),
						E: expr.Raw("test_query"),
					},
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
					Query: litsql.QueryFunc{
						D: testutils.NewTestDialect(),
						E: expr.Raw("test_query2"),
					},
				},
			},
		})
	assert.Assert(t, len(clause.CTEs) == 2)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WITH testCTE(id, name) AS test_query,")
	o.WriteSeparator()
	o.Write("testCTE2(id, name) AS test_query2")
	testutils.TestExpression(t, clause, o)
}
