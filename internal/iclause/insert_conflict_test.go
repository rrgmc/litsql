package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestInsertConflict(t *testing.T) {
	clause := &InsertConflict{
		Do: "UPDATE",
		Set: Set{
			Set: []litsql.Expression{
				expr.Raw("id = 1"),
				expr.Raw("age = 50"),
			},
		},
		Where: Where{
			Conditions: []litsql.Expression{
				expr.Raw("id = 5"),
				expr.Raw("age = 10"),
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("ON CONFLICT DO UPDATE")

	o.WriteSeparator()
	o.Write("SET")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("id = 1,")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("age = 50")

	o.WriteSeparator()
	o.Write("WHERE id = 5 AND age = 10")

	testutils.TestExpression(t, clause, o)
}

func TestInsertConflictEmpty(t *testing.T) {
	clause := &InsertConflict{}
	testutils.TestExpressionErrorIs(t, clause, litsql.ErrClause)
}

func TestInsertConflictTarget(t *testing.T) {
	clause := &InsertConflict{
		Do: "UPDATE",
		Target: InsertConflictTarget{
			Where: []litsql.Expression{
				expr.Raw("id = 5"),
				expr.Raw("age = 10"),
			},
		},
		Set: Set{
			Set: []litsql.Expression{
				expr.Raw("id = 1"),
				expr.Raw("age = 50"),
			},
		},
		Where: Where{
			Conditions: []litsql.Expression{
				expr.Raw("id = 5"),
				expr.Raw("age = 10"),
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("ON CONFLICT ")
	o.Write("WHERE id = 5 AND age = 10 ")
	o.Write("DO UPDATE")

	o.WriteSeparator()
	o.Write("SET")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("id = 1,")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("age = 50")

	o.WriteSeparator()
	o.Write("WHERE id = 5 AND age = 10")

	testutils.TestExpression(t, clause, o)
}
