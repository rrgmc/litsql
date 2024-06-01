package psql_test

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/dm"
	"gotest.tools/v3/assert"
)

func TestDelete(t *testing.T) {
	query := psql.Delete(
		dm.From("user"),
		dm.Where("id = 1"),
	)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 0)
}
