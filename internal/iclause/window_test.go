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
	testutils.TestWriterExpression(t, clause, o)
}

func TestWindowEmpty(t *testing.T) {
	clause := &Windows{}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestWriterExpression(t, clause, o)
}

func TestWindowClauses(t *testing.T) {
	clause := &Windows{
		Windows: []*NamedWindow{
			{
				Name: "window_test",
				Definition: WindowDef{
					From: "other_window",
					PartitionBy: []litsql.Expression{
						expr.Raw("id"),
						expr.Raw("name"),
					},
					OrderBy: []litsql.Expression{
						expr.Raw("age"),
						expr.Raw("last_activity_date"),
					},
				},
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WINDOW window_test AS (other_window PARTITION BY id, name ORDER BY age, last_activity_date)")
	testutils.TestWriterExpression(t, clause, o)
}

func TestWindowClausesSpacing1(t *testing.T) {
	clause := &Windows{
		Windows: []*NamedWindow{
			{
				Name: "window_test",
				Definition: WindowDef{
					From: "other_window",
					OrderBy: []litsql.Expression{
						expr.Raw("age"),
						expr.Raw("last_activity_date"),
					},
				},
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WINDOW window_test AS (other_window ORDER BY age, last_activity_date)")
	testutils.TestWriterExpression(t, clause, o)
}

func TestWindowClausesSpacing2(t *testing.T) {
	clause := &Windows{
		Windows: []*NamedWindow{
			{
				Name: "window_test",
				Definition: WindowDef{
					OrderBy: []litsql.Expression{
						expr.Raw("age"),
						expr.Raw("last_activity_date"),
					},
				},
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WINDOW window_test AS (ORDER BY age, last_activity_date)")
	testutils.TestWriterExpression(t, clause, o)
}

func TestWindowFrame(t *testing.T) {
	clause := &Windows{
		Windows: []*NamedWindow{
			{
				Name: "window_test",
				Definition: WindowDef{
					Frame: Frame{
						Defined:   true,
						Mode:      "ROWS",
						Start:     expr.Raw("UNBOUNDED PRECEDING"),
						End:       expr.Raw("CURRENT ROW"),
						Exclusion: "EXCLUDE GROUP",
					},
				},
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("WINDOW window_test AS (ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW EXCLUDE EXCLUDE GROUP)")
	testutils.TestWriterExpression(t, clause, o)
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
	testutils.TestWriterExpression(t, clause, o)
}
