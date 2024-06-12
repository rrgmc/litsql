package ichain

import (
	"testing"

	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestWindowChain(t *testing.T) {
	chain := NewWindowChain[testutils.TestTag, Window[testutils.TestTag]](&WindowChain[testutils.TestTag, Window[testutils.TestTag]]{
		Windows:     &iclause.Windows{},
		NamedWindow: &iclause.NamedWindow{},
	})

	chain.
		From("from_test").
		PartitionBy("p1").
		OrderBy("o1").
		Frame(expr.Raw("ROWS UNBOUNDED PRECEDING"))

	assert.Equal(t, "from_test", chain.NamedWindow.Definition.From)
	assert.Assert(t, len(chain.NamedWindow.Definition.PartitionBy) == 1)
	testutils.TestExpression(t, chain.NamedWindow.Definition.PartitionBy[0], "p1")
	assert.Assert(t, len(chain.NamedWindow.Definition.OrderBy) == 1)
	testutils.TestExpression(t, chain.NamedWindow.Definition.OrderBy[0], "o1")
	testutils.TestExpression(t, chain.NamedWindow.Definition.Frame, "ROWS UNBOUNDED PRECEDING")
}
