package iclause

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestDeleteOnly(t *testing.T) {
	c := &DeleteOnly{
		Only: true,
	}

	o := testutils.NewTestBuffer()
	o.Write("ONLY")
	testutils.TestExpression(t, c, o)
}

func TestDeleteOnlyFalse(t *testing.T) {
	c := &DeleteOnly{
		Only: false,
	}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestExpression(t, c, o)
}
