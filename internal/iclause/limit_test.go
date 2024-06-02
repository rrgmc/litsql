package iclause

import (
	"testing"

	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestLimit(t *testing.T) {
	clause := &Limit{
		Count: expr.Raw("10"),
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("LIMIT 10")
	testutils.TestExpression(t, clause, o)
}

func TestLimitEmpty(t *testing.T) {
	clause := &Limit{}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestExpression(t, clause, o)
}
