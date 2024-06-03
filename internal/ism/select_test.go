package ism

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestSelect(t *testing.T) {
	expectedQuery := "SELECT id, name FROM users WHERE age < 10"
	var expectedArgs []any

	query := Select[testutils.TestTag](testutils.NewTestDialect(),
		Columns[testutils.TestTag]("id", "name"),
		From[testutils.TestTag]("users"),
		Where[testutils.TestTag]("age < 10"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
