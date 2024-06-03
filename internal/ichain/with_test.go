package ichain

import (
	"testing"

	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestWithChain(t *testing.T) {
	chain := &WithChain[testutils.TestTag]{
		With: &iclause.With{},
		CTE:  &iclause.CTE{},
	}

	chain.
		Recursive().
		NotMaterialized().
		SearchDepth("scol1", "stcol1", "stcol2").
		Cycle("ccol1", "cusing1", "ctcol1", "ctcol2").
		CycleValue(15, 20)

	assert.Assert(t, chain.With.Recursive != nil && *chain.With.Recursive)
	assert.Assert(t, chain.CTE.Materialized != nil && *chain.CTE.Materialized == false)
	assert.Equal(t, "scol1", chain.CTE.Search.Set)
	assert.DeepEqual(t, []string{"stcol1", "stcol2"}, chain.CTE.Search.Columns)
	assert.Equal(t, "ccol1", chain.CTE.Cycle.Set)
	assert.Equal(t, "cusing1", chain.CTE.Cycle.Using)
	assert.DeepEqual(t, []string{"ctcol1", "ctcol2"}, chain.CTE.Cycle.Columns)
	testutils.TestExpression(t, chain.CTE.Cycle.SetVal, "$1", 15)
	testutils.TestExpression(t, chain.CTE.Cycle.DefaultVal, "$1", 20)
}
