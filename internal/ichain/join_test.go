package ichain

import (
	"testing"

	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestJoinChain(t *testing.T) {
	chain := NewJoinChain[testutils.TestTag, Join[testutils.TestTag]](&JoinChain[testutils.TestTag, Join[testutils.TestTag]]{
		Join: &iclause.Join{
			To: &iclause.From{},
		},
	})

	chain.
		As("test_alias", "c1", "c2").
		Only().
		Lateral().
		WithOrdinality().
		On("on_str").
		Using("using_str").
		Natural()

	assert.Equal(t, "test_alias", chain.Join.To.Alias)
	assert.DeepEqual(t, []string{"c1", "c2"}, chain.Join.To.AliasColumns)
	assert.Equal(t, true, chain.Join.To.Only)
	assert.Equal(t, true, chain.Join.To.Lateral)
	assert.Equal(t, true, chain.Join.To.WithOrdinality)
	assert.DeepEqual(t, []string{"using_str"}, chain.Join.Using)
	assert.Assert(t, len(chain.Join.On) == 1)
	testutils.TestExpression(t, chain.Join.On[0], "on_str")
}
