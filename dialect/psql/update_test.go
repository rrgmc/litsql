package psql_test

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/um"
	"gotest.tools/v3/assert"
)

func TestUpdate(t *testing.T) {
	query := psql.Update(
		um.Table("user"),
		um.Set("name", "joao"),
		um.Where("id = 1"),
	)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 1)
}

func TestUpdateRaw(t *testing.T) {
	query := psql.UpdateRaw("update users set name = ? where id = ?", 12, 15)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Equal(t, "update users set name = ? where id = ?", queryStr)
	assert.DeepEqual(t, []any{12, 15}, params)
}
