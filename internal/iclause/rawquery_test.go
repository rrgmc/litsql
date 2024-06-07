package iclause

import (
	"testing"

	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestRawQuery(t *testing.T) {
	clause := &RawQuery{
		Query: expr.Raw("SELECT * FROM users"),
		Args:  []any{1, 5},
	}

	o := testutils.NewTestBuffer()
	o.Write("SELECT * FROM users")
	testutils.TestWriterExpression(t, clause, o, 1, 5)
}

func TestRawQueryExtraArgs(t *testing.T) {
	clause := &RawQuery{
		Query: expr.Clause("SELECT * FROM users WHERE id = ?", 66),
		Args:  []any{1, 5},
	}

	o := testutils.NewTestBuffer()
	o.Write("SELECT * FROM users WHERE id = $1")
	testutils.TestWriterExpression(t, clause, o, 66, 1, 5)
}

func TestRawQueryEmpty(t *testing.T) {
	clause := &RawQuery{}

	o := testutils.NewTestBuffer()
	testutils.TestWriterExpression(t, clause, o)
}
