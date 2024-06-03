package ichain

import (
	"testing"

	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestWindowChain(t *testing.T) {
	chain := &WindowChain[testutils.TestTag]{
		Windows:     &iclause.Windows{},
		NamedWindow: &iclause.NamedWindow{},
	}

	chain.
		From("from_test").
		PartitionBy("p1").
		OrderBy("o1").
		Rows().
		FromFollowing(expr.Raw("ff")).
		ToPreceding(expr.Raw("pc")).
		ExcludeGroup()

	assert.Equal(t, "from_test", chain.NamedWindow.Definition.From)
	assert.Assert(t, len(chain.NamedWindow.Definition.PartitionBy) == 1)
	testutils.TestExpression(t, chain.NamedWindow.Definition.PartitionBy[0], "p1")
	assert.Assert(t, len(chain.NamedWindow.Definition.OrderBy) == 1)
	testutils.TestExpression(t, chain.NamedWindow.Definition.OrderBy[0], "o1")
	assert.Equal(t, "ROWS", chain.NamedWindow.Definition.Mode)
	testutils.TestExpression(t, chain.NamedWindow.Definition.Start, "ff FOLLOWING")
	testutils.TestExpression(t, chain.NamedWindow.Definition.End, "pc PRECEDING")
	assert.Equal(t, "GROUP", chain.NamedWindow.Definition.Exclusion)
}
