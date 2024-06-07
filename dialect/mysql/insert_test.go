package mysql_test

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/mysql"
	"github.com/rrgmc/litsql/dialect/mysql/im"
	"github.com/rrgmc/litsql/expr"
	"gotest.tools/v3/assert"
)

func TestInsert(t *testing.T) {
	query := mysql.Insert(
		im.Into("device", "id", "name"),
		im.Values(expr.String("d1")),
	)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 0)
}
