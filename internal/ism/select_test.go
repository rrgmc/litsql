package ism

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
	"github.com/rrgmc/litsql/sq"
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

func TestSelectBasic(t *testing.T) {
	expectedQuery := "SELECT DISTINCT u.id, u.name, ua.address FROM users AS u INNER JOIN users_address AS a ON users.id = a.user_id WHERE age < 10 ORDER BY name ASC OFFSET 10 LIMIT 100"
	var expectedArgs []any

	query := Select[testutils.TestTag](testutils.NewTestDialect(),
		Distinct[testutils.TestTag](),
		Columns[testutils.TestTag]("u.id", "u.name", "ua.address"),
		From[testutils.TestTag]("users AS u"),
		InnerJoin[testutils.TestTag]("users_address AS a").On("users.id = a.user_id"),
		Where[testutils.TestTag]("age < 10"),
		OrderBy[testutils.TestTag]("name ASC"),
		Offset[testutils.TestTag](10),
		Limit[testutils.TestTag](100),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectBasicArgs(t *testing.T) {
	expectedQuery := "SELECT DISTINCT u.id, u.name, ua.address FROM users AS u INNER JOIN users_address AS a ON users.id = a.user_id WHERE age < $1 ORDER BY name ASC OFFSET $2 LIMIT $3"
	expectedArgs := []any{10, 10, 100}

	query := Select[testutils.TestTag](testutils.NewTestDialect(),
		Distinct[testutils.TestTag](),
		Columns[testutils.TestTag]("u.id", "u.name", "ua.address"),
		From[testutils.TestTag]("users AS u"),
		InnerJoin[testutils.TestTag]("users_address AS a").On("users.id = a.user_id"),
		WhereC[testutils.TestTag]("age < ?", 10),
		OrderBy[testutils.TestTag]("name ASC"),
		OffsetA[testutils.TestTag](10),
		LimitA[testutils.TestTag](100),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectBasicNamedArgs(t *testing.T) {
	expectedQuery := "SELECT DISTINCT u.id, u.name, ua.address FROM users AS u INNER JOIN users_address AS a ON users.id = a.user_id WHERE age < $1 ORDER BY name ASC OFFSET $2 LIMIT $3"
	expectedArgs := []any{10, 10, 100}

	query := Select[testutils.TestTag](testutils.NewTestDialect(),
		Distinct[testutils.TestTag](),
		Columns[testutils.TestTag]("u.id", "u.name", "ua.address"),
		From[testutils.TestTag]("users AS u"),
		InnerJoin[testutils.TestTag]("users_address AS a").On("users.id = a.user_id"),
		WhereC[testutils.TestTag]("age < ?", sq.Arg("age")),
		OrderBy[testutils.TestTag]("name ASC"),
		OffsetA[testutils.TestTag](sq.Arg("offset")),
		LimitA[testutils.TestTag](sq.Arg("limit")),
	)

	testutils.TestQueryParseArgs(t, query, expectedQuery, map[string]any{
		"age":    10,
		"offset": 10,
		"limit":  100,
	}, expectedArgs...)
}
