package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestWindow(t *testing.T) {
	clause := &Windows{
		Windows: []*NamedWindow{
			{
				Name: "window_test",
				Definition: WindowDef{
					PartitionBy: []litsql.Expression{
						expr.Raw("id"),
						expr.Raw("name"),
					},
				},
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WINDOW window_test AS (PARTITION BY id, name)")
	testutils.TestExpression(t, clause, o)
}

func TestWindowEmpty(t *testing.T) {
	clause := &Windows{}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestExpression(t, clause, o)
}

func TestWindowMerge(t *testing.T) {
	clause := testutils.Merge(
		&Windows{
			Windows: []*NamedWindow{
				{
					Name: "window_test",
					Definition: WindowDef{
						PartitionBy: []litsql.Expression{
							expr.Raw("id"),
							expr.Raw("name"),
						},
					},
				},
			},
		},
		&Windows{
			Windows: []*NamedWindow{
				{
					Name: "window_test2",
					Definition: WindowDef{
						PartitionBy: []litsql.Expression{
							expr.Raw("id"),
							expr.Raw("name"),
						},
					},
				},
			},
		})
	assert.Assert(t, len(clause.Windows) == 2)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WINDOW")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("window_test AS (PARTITION BY id, name),")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("window_test2 AS (PARTITION BY id, name)")
	testutils.TestExpression(t, clause, o)
}
