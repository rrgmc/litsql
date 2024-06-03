package ism

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestSelect(t *testing.T) {
	expectedQuery := "SELECT id, name FROM users"
	var expectedArgs []any

	query := Select[testutils.TestTag](testutils.NewTestDialect(),
		Columns[testutils.TestTag]("id", "name"),
		From[testutils.TestTag]("users"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
