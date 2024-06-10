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
		InsertConflictUpdate: &iclause.InsertConflictUpdate{},
	}

	chain.
		Where("a = 1").
		DoUpdate(
			mod.InsertConflictUpdateModFunc[testutils.TestTag, imod.InsertConflictUpdateModTag](func(a *iclause.InsertConflictUpdate) {
				a.Set.Set = append(a.Set.Set, expr.Raw("a = 1"))
			}),
		)

	assert.Assert(t, len(chain.InsertConflictUpdate.Target.Where) == 1)
	testutils.TestExpression(t, chain.InsertConflictUpdate.Target.Where[0], "a = 1")
	assert.Equal(t, "UPDATE", chain.InsertConflictUpdate.Do)
	assert.Assert(t, len(chain.InsertConflictUpdate.Set.Set) == 1)
	testutils.TestExpression(t, chain.InsertConflictUpdate.Set.Set[0], "a = 1")
}
