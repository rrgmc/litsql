package iclause

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestInsertOverriding(t *testing.T) {
	clause := &InsertOverriding{
		Overriding: "SYSTEM",
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("OVERRIDING SYSTEM VALUE")
	testutils.TestWriterExpression(t, clause, o)
}

func TestInsertOverridingEmpty(t *testing.T) {
	clause := &InsertOverriding{}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestWriterExpression(t, clause, o)
}
