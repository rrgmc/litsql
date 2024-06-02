package sq

import (
	"database/sql"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestParseArgValues(t *testing.T) {
	args := Args{
		1,
		Arg("first"),
		ArgDefault("second", 55),
		DBArg("third"),
		DBArgDefault("fourth", 77),
		ArgFunc(func() (any, error) {
			return "in-func", nil
		}),
	}
	pargs, err := args.Parse(map[string]any{
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
