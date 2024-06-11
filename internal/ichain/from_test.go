package ichain

import (
	"testing"

	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestFromChain(t *testing.T) {
	chain := NewFromChain[testutils.TestTag, From[testutils.TestTag]](&FromChain[testutils.TestTag, From[testutils.TestTag]]{
		From: &iclause.From{},
	})

	chain.
		As("test_alias", "c1", "c2").
		Only().
		Lateral().
		WithOrdinality()

	assert.Equal(t, "test_alias", chain.From.Alias)
	assert.DeepEqual(t, []string{"c1", "c2"}, chain.From.AliasColumns)
	assert.Equal(t, true, chain.From.Only)
	assert.Equal(t, true, chain.From.Lateral)
	assert.Equal(t, true, chain.From.WithOrdinality)
}
