package iclause

import (
	"testing"

	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestOffset(t *testing.T) {
	clause := &Offset{
		Count: expr.Raw("10"),
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("OFFSET 10")
	testutils.TestWriterExpression(t, clause, o)
}

func TestOffsetEmpty(t *testing.T) {
	clause := &Offset{}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestWriterExpression(t, clause, o)
}
