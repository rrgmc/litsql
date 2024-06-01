package mysql_test

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/mysql"
	"github.com/rrgmc/litsql/dialect/mysql/um"
	"gotest.tools/v3/assert"
)

func TestUpdate(t *testing.T) {
	query := mysql.Update(
		um.Table("user"),
		um.Set("name", "joao"),
		um.Where("id = 1"),
	)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 1)
}
