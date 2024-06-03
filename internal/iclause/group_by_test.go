package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestGroupBy(t *testing.T) {
	clause := &GroupBy{
		Groups: []litsql.Expression{
			expr.Raw("id"),
			expr.Raw("name"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("GROUP BY id, name")
	testutils.TestWriterExpression(t, clause, o)
}

func TestGroupByEmpty(t *testing.T) {
	clause := &GroupBy{}

	o := testutils.NewTestBuffer()
	testutils.TestWriterExpression(t, clause, o)
}

func TestGroupByDistinct(t *testing.T) {
	clause := &GroupBy{
		Groups: []litsql.Expression{
			expr.Raw("id"),
			expr.Raw("name"),
		},
		Distinct: true,
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("GROUP BY DISTINCT id, name")
	testutils.TestWriterExpression(t, clause, o)
}

func TestGroupByWith(t *testing.T) {
	clause := &GroupBy{
		Groups: []litsql.Expression{
			expr.Raw("id"),
			expr.Raw("name"),
		},
		With: "ROLLUP",
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("GROUP BY id, name WITH ROLLUP")
	testutils.TestWriterExpression(t, clause, o)
}

func TestGroupByMerge(t *testing.T) {
	clause, err := internal.MergeClauses(
		&GroupBy{
			Groups: []litsql.Expression{expr.Raw("id"), expr.Raw("id2")},
		},
		&GroupBy{
			Groups: []litsql.Expression{expr.Raw("id3"), expr.Raw("id4")},
		})
	assert.NilError(t, err)
	assert.Assert(t, len(clause.Groups) == 4)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("GROUP BY id, id2, id3, id4")
	testutils.TestWriterExpression(t, clause, o)
}
