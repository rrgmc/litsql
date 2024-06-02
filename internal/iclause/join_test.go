package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestJoin(t *testing.T) {
	clause := &Join{
		Type: "INNER JOIN",
		To: &From{
			Table: expr.Raw("users"),
		},
		On: []litsql.Expression{
			expr.Raw("id = 5"),
			expr.Raw("age = 10"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("INNER JOIN users ON id = 5 AND age = 10")
	testutils.TestWriterExpression(t, clause, o)
}

func TestJoinEmpty(t *testing.T) {
	clause := &Join{}
	testutils.TestWriterExpressionErrorIs(t, clause, litsql.ErrClause)
}

func TestJoinEmptyFrom(t *testing.T) {
	clause := &Join{
		Type: "INNER JOIN",
		To:   &From{},
	}
	testutils.TestWriterExpressionErrorIs(t, clause, litsql.ErrClause)
}

func TestJoinEmptyOn(t *testing.T) {
	clause := &Join{
		Type: "INNER JOIN",
		To: &From{
			Table: expr.Raw("users"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("INNER JOIN users")
	testutils.TestWriterExpression(t, clause, o)
}

func TestJoinFlags(t *testing.T) {
	clause := &Join{
		Type: "INNER JOIN",
		To: &From{
			Table: expr.Raw("users"),
		},
		On: []litsql.Expression{
			expr.Raw("id = 5"),
			expr.Raw("age = 10"),
		},
		Natural: true,
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("NATURAL INNER JOIN users ON id = 5 AND age = 10")
	testutils.TestWriterExpression(t, clause, o)
}

func TestJoinUsing(t *testing.T) {
	clause := &Join{
		Type: "INNER JOIN",
		To: &From{
			Table: expr.Raw("users"),
		},
		On: []litsql.Expression{
			expr.Raw("id = 5"),
			expr.Raw("age = 10"),
		},
		Using: []string{"id", "age"},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("INNER JOIN users ON id = 5 AND age = 10 USING(id, age)")
	testutils.TestWriterExpression(t, clause, o)
}
