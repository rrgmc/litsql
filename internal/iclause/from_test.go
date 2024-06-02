package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestFrom(t *testing.T) {
	clause := &From{
		Table:   expr.Raw("users"),
		Starter: true,
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("FROM users")
	testutils.TestWriterExpression(t, clause, o)
}

func TestFromEmpty(t *testing.T) {
	clause := &From{}
	testutils.TestWriterExpressionErrorIs(t, clause, litsql.ErrClause)
}

func TestFromNonStarter(t *testing.T) {
	clause := &From{
		Table:   expr.Raw("users"),
		Starter: false,
	}

	o := testutils.NewTestBuffer()
	o.Write("users")
	testutils.TestWriterExpression(t, clause, o)
}

func TestFromClause(t *testing.T) {
	clause := &From{
		Table:   expr.Raw("users"),
		Starter: true,
		Clause:  "USING",
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("USING users")
	testutils.TestWriterExpression(t, clause, o)
}

func TestFromFlags(t *testing.T) {
	clause := &From{
		Table:          expr.Raw("users"),
		Starter:        true,
		Alias:          "u",
		Only:           true,
		Lateral:        true,
		WithOrdinality: true,
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("FROM ONLY LATERAL users WITH ORDINALITY AS u")
	testutils.TestWriterExpression(t, clause, o)
}

func TestFromColumns(t *testing.T) {
	clause := &From{
		Table:        expr.Raw("users"),
		Starter:      true,
		Alias:        "u",
		AliasColumns: []string{"id", "name"},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("FROM users AS u(id, name)")
	testutils.TestWriterExpression(t, clause, o)
}
