package argstruct

import (
	"testing"

	"github.com/rrgmc/litsql/sq"
	"gotest.tools/v3/assert"
)

func TestArgStruct(t *testing.T) {
	type x struct {
		H string
		J int
		L int `r:"LA"`
		M int `r:"MM,omitempty,x=15"`
		O *int
		P *string
	}

	oval := 45

	value := &x{
		H: "99",
		J: 11,
		L: 45,
		M: 91,
		O: &oval,
	}

	args := []any{
		81,
		sq.NamedArg("H"),
		sq.NamedArg("J"),
		sq.NamedArg("LA"),
		sq.NamedArg("MM"),
		sq.NamedArg("O"),
		sq.NamedArg("P"),
	}
	pargs, err := sq.ParseArgs(args, Values(value, WithTagName("r")))
	assert.NilError(t, err)

	assert.DeepEqual(t, []any{
		81,
		"99",
		11,
		45,
		91,
		&oval,
		nil,
	}, pargs)
}
