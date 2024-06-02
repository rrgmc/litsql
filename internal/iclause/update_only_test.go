package iclause

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestUpdateOnly(t *testing.T) {
	c := &UpdateOnly{
		Only: true,
	}

	o := testutils.NewTestBuffer()
	o.Write("ONLY")
	testutils.TestWriterExpression(t, c, o)
}

func TestUpdateOnlyFalse(t *testing.T) {
	c := &UpdateOnly{
		Only: false,
	}

	o := testutils.NewTestBuffer()
	o.Write("")
	testutils.TestWriterExpression(t, c, o)
}
