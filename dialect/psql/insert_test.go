package psql_test

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/im"
	"gotest.tools/v3/assert"
)

func TestInsert(t *testing.T) {
	query := psql.Insert(
		im.Into("device", "id", "name"),
		im.ValuesS("d1"),
	)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 0)
}
