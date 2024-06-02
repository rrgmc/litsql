package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestCombine(t *testing.T) {
	clause := &Combine{
		Strategy: "UNION",
		Query: litsql.QueryFunc{
			D: testutils.NewTestDialect(),
			E: expr.Raw("test_query"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("UNION test_query")
	testutils.TestExpression(t, clause, o)
}

func TestCombineAll(t *testing.T) {
	clause := &Combine{
		Strategy: "UNION",
		All:      true,
		Query: litsql.QueryFunc{
			D: testutils.NewTestDialect(),
			E: expr.Raw("test_query"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("UNION ALL test_query")
	testutils.TestExpression(t, clause, o)
}

func TestCombineEmpty(t *testing.T) {
	clause := &Combine{}
	testutils.TestExpressionErrorIs(t, clause, litsql.ErrClause)
}
