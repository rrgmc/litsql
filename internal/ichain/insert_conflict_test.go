package ichain

import (
	"testing"

	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/imod"
	"github.com/rrgmc/litsql/internal/testutils"
	"github.com/rrgmc/litsql/sq/mod"
	"gotest.tools/v3/assert"
)

func TestInsertConflictChain(t *testing.T) {
	chain := &InsertConflictChain[testutils.TestTag]{
		InsertConflict: &iclause.InsertConflict{},
	}

	chain.
		Where("a = 1").
		DoUpdate(
			mod.InsertConflictUpdateModFunc[testutils.TestTag, imod.InsertConflictUpdateModUM](func(a *iclause.InsertConflict) {
				a.Set.Set = append(a.Set.Set, expr.Raw("a = 1"))
			}),
		)

	assert.Assert(t, len(chain.InsertConflict.Target.Where) == 1)
	testutils.TestExpression(t, chain.InsertConflict.Target.Where[0], "a = 1")
	assert.Equal(t, "UPDATE", chain.InsertConflict.Do)
	assert.Assert(t, len(chain.InsertConflict.Set.Set) == 1)
	testutils.TestExpression(t, chain.InsertConflict.Set.Set[0], "a = 1")
}
