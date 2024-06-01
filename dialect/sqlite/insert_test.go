package sqlite_test

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/sqlite"
	"github.com/rrgmc/litsql/dialect/sqlite/im"
	"github.com/rrgmc/litsql/expr"
	"gotest.tools/v3/assert"
)

func TestInsert(t *testing.T) {
	query := sqlite.Insert(
		im.Into("device", "id", "name"),
		im.Values(expr.S("d1")),
	)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 0)
}
