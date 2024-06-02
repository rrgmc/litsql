package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
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

// func TestWindowMerge(t *testing.T) {
// 	clause := testutils.Merge(
// 		&Windows{
// 			Windows: []litsql.Expression{expr.Raw("id"), expr.Raw("id2")},
// 		},
// 		&Windows{
// 			Windows: []litsql.Expression{expr.Raw("id3"), expr.Raw("id4")},
// 		})
// 	assert.Assert(t, len(clause.Windows) == 2)
//
// 	o := testutils.NewTestBuffer()
// 	o.Write("id, id2, id3, id4")
// 	testutils.TestExpression(t, clause, o)
// }
