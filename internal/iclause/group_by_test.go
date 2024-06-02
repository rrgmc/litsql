package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
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
	testutils.TestExpression(t, clause, o)
}

func TestGroupByEmpty(t *testing.T) {
	clause := &GroupBy{}

	o := testutils.NewTestBuffer()
	testutils.TestExpression(t, clause, o)
}

func TestGroupByMerge(t *testing.T) {
	clause := testutils.Merge(
		&GroupBy{
			Groups: []litsql.Expression{expr.Raw("id"), expr.Raw("id2")},
		},
		&GroupBy{
			Groups: []litsql.Expression{expr.Raw("id3"), expr.Raw("id4")},
		})
	assert.Assert(t, len(clause.Groups) == 4)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("GROUP BY id, id2, id3, id4")
	testutils.TestExpression(t, clause, o)
}
