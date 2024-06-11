package ichain

import (
	"testing"

	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestGroupByChain(t *testing.T) {
	chain := NewGroupByChain[testutils.TestTag, GroupBy[testutils.TestTag]](&GroupByChain[testutils.TestTag, GroupBy[testutils.TestTag]]{
		GroupBy: &iclause.GroupBy{},
	})

	// chain := &GroupByChain[testutils.TestTag]{
	// 	GroupBy: &iclause.GroupBy{},
	// }

	chain.
		Distinct().
		With("test_with")

	assert.Equal(t, true, chain.GroupBy.Distinct)
	assert.Equal(t, "test_with", chain.GroupBy.With)
}
