package mysql_test

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/mysql"
	"github.com/rrgmc/litsql/dialect/mysql/sm"
	"gotest.tools/v3/assert"
)

func TestSelect(t *testing.T) {
	query := mysql.Select(
		sm.Columns("a", "b"),
		sm.From("device"),
	)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 0)
}
