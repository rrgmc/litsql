package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestDistinct(t *testing.T) {
	c := &Distinct{
		On: []litsql.Expression{
			expr.Raw("id"),
			expr.Raw("name"),
		},
	}

	o := testutils.NewTestBuffer()
	o.Write("DISTINCT ON (id, name)")
	testutils.TestExpression(t, c, o)
}

func TestDistinctEmpty(t *testing.T) {
	clause := &Distinct{}

	o := testutils.NewTestBuffer()
	o.Write("DISTINCT")
	testutils.TestExpression(t, clause, o)
}

func TestDistinctMerge(t *testing.T) {
	clause := testutils.Merge(
		&Distinct{
			On: []litsql.Expression{expr.Raw("id"), expr.Raw("id2")},
		},
		&Distinct{
			On: []litsql.Expression{expr.Raw("id3"), expr.Raw("id4")},
		})
	assert.Assert(t, len(clause.On) == 4)

	o := testutils.NewTestBuffer()
	o.Write("DISTINCT ON (id, id2, id3, id4)")
	testutils.TestExpression(t, clause, o)
}
