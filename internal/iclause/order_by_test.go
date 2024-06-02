package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestOrderBy(t *testing.T) {
	clause := &OrderBy{
		Expressions: []litsql.Expression{
			expr.Raw("id ASC"),
			expr.Raw("name DESC"),
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("ORDER BY id ASC, name DESC")
	testutils.TestWriterExpression(t, clause, o)
}

func TestOrderByEmpty(t *testing.T) {
	clause := &OrderBy{}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestWriterExpression(t, clause, o)
}

func TestOrderByMerge(t *testing.T) {
	clause := testutils.Merge(
		&OrderBy{
			Expressions: []litsql.Expression{expr.Raw("id"), expr.Raw("id2")},
		},
		&OrderBy{
			Expressions: []litsql.Expression{expr.Raw("id3"), expr.Raw("id4")},
		})
	assert.Assert(t, len(clause.Expressions) == 4)

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("ORDER BY id, id2, id3, id4")
	testutils.TestWriterExpression(t, clause, o)
}
