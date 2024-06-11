package psql_test

import (
	"fmt"
	"testing"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"gotest.tools/v3/assert"
)

func TestSelect(t *testing.T) {
	query := psql.Select(
		sm.Columns("a", "b"),
		sm.From("device"),
	)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 0)
}

func TestSelectRaw(t *testing.T) {
	query := psql.SelectRaw("select * from users where id = ?", 12, 15)
	queryStr, params, err := query.Build()

	assert.NilError(t, err)
	assert.Equal(t, "select * from users where id = ?", queryStr)
	assert.DeepEqual(t, []any{12, 15}, params)
}

func TestSelectG(t *testing.T) {
	query := psql.Select(
		sm.Columns("a", "b"),
		sm.From("device"),
		sm.GroupBy("b").Distinct(),
	)
	queryStr, params, err := query.Build()

	fmt.Println(queryStr)

	assert.NilError(t, err)
	assert.Assert(t, queryStr != "")
	assert.Assert(t, len(params) == 0)
}
