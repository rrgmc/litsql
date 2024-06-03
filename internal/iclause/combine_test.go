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
		Query:    litsql.QueryFunc(testutils.NewTestDialect(), expr.Raw("test_query"), nil),
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("UNION test_query")
	testutils.TestWriterExpression(t, clause, o)
}

func TestCombineAll(t *testing.T) {
	clause := &Combine{
		Strategy: "UNION",
		All:      true,
		Query:    litsql.QueryFunc(testutils.NewTestDialect(), expr.Raw("test_query"), nil),
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("UNION ALL test_query")
	testutils.TestWriterExpression(t, clause, o)
}

func TestCombineEmpty(t *testing.T) {
	clause := &Combine{}
	testutils.TestWriterExpressionErrorIs(t, clause, litsql.ErrClause)
}
