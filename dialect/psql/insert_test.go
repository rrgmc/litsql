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
		im.ValuesString("d1"),
	)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 0)
}

func TestInsertRaw(t *testing.T) {
	query := psql.InsertRaw("insert into users (id) values (?)", 12, 15)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Equal(t, "insert into users (id) values (?)", queryStr)
	assert.DeepEqual(t, []any{12, 15}, params)
}
