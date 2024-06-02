package iclause

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestDeleteFrom(t *testing.T) {
	c := &DeleteFrom{
		Table: "users",
	}

	o := testutils.NewTestBuffer()
	o.Write("users")
	testutils.TestExpression(t, c, o)
}
