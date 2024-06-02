package iclause

import (
	"testing"

	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestTable(t *testing.T) {
	clause := &Table{
		Expression: expr.Raw("users"),
	}

	o := testutils.NewTestBuffer()
	o.Write("users")
	testutils.TestExpression(t, clause, o)
}

func TestTableEmpty(t *testing.T) {
	clause := &Table{}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestExpression(t, clause, o)
}

func TestTableColumns(t *testing.T) {
	clause := &Table{
		Expression: expr.Raw("users"),
		Columns:    []string{"id", "name"},
	}

	o := testutils.NewTestBuffer()
	o.Write("users (id, name)")
	testutils.TestExpression(t, clause, o)
}

func TestTableAlias(t *testing.T) {
	clause := &Table{
		Expression: expr.Raw("users"),
		Alias:      "u",
		Columns:    []string{"id", "name"},
	}

	o := testutils.NewTestBuffer()
	o.Write("users AS u (id, name)")
	testutils.TestExpression(t, clause, o)
}
