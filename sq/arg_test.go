package sq

import (
	"database/sql"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rrgmc/litsql"
	"gotest.tools/v3/assert"
)

func TestParseArgValues(t *testing.T) {
	args := []any{
		1,
		NamedArg("first"),
		NamedArg("second", WithDefaultValue(55)),
		DBNamedArg("third"),
		DBNamedArg("fourth", WithDefaultValue(77)),
		ArgFunc(func() (any, error) {
			return "in-func", nil
		}),
	}
	pargs, err := ParseArgs(args, litsql.MapArgValues{
		"first": 99,
		"third": 45,
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, []any{
		1,
		99,
		55,
		sql.Named("third", 45),
		sql.Named("fourth", 77),
		"in-func",
	}, pargs,
		cmp.Comparer(func(x, y sql.NamedArg) bool {
			return x.Name == y.Name && cmp.Equal(x.Value, y.Value)
		}))
}
