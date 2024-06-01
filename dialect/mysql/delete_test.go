package mysql_test

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/mysql"
	"github.com/rrgmc/litsql/dialect/mysql/dm"
	"gotest.tools/v3/assert"
)

func TestDelete(t *testing.T) {
	query := mysql.Delete(
		dm.From("user"),
		dm.Where("id = 1"),
	)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 0)
}
