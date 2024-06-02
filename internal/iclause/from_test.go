package iclause

import (
	"testing"

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
	testutils.TestExpression(t, clause, o)
}

func TestFromEmpty(t *testing.T) {
	clause := &From{}
	testutils.TestExpressionIsError(t, clause)
}

func TestFromNonStarter(t *testing.T) {
	clause := &From{
		Table:   expr.Raw("users"),
		Starter: false,
	}

	o := testutils.NewTestBuffer()
	o.Write("users")
	testutils.TestExpression(t, clause, o)
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
	testutils.TestExpression(t, clause, o)
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
	testutils.TestExpression(t, clause, o)
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
	testutils.TestExpression(t, clause, o)
}
