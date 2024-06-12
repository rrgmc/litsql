package im

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/mysql"
	"github.com/rrgmc/litsql/dialect/mysql/sm"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestInsert(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES (?, ?)"
	expectedArgs := []any{15, "John"}

	query := mysql.Insert(
		Into("users", "id", "name"),
		Values(15, "John"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertBasic(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES (?, ?), (?, ?) RETURNING id"
	expectedArgs := []any{15, "John", 16, "Mary"}

	query := mysql.Insert(
		Into("users", "id", "name"),
		Values(15, "John"),
		Values(16, "Mary"),
		Returning("id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertOverriding(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES (?, ?)"
	expectedArgs := []any{15, "John"}

	query := mysql.Insert(
		Into("users", "id", "name"),
		Values(15, "John"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertWith(t *testing.T) {
	expectedQuery := "WITH city(city_id) AS (SELECT city FROM users WHERE id = ?) INSERT INTO users (id, name) VALUES (?, ?)"
	expectedArgs := []any{2, 15, "John"}

	query := mysql.Insert(
		With("city", "city_id").As(
			mysql.Select(
				sm.Columns("city"),
				sm.From("users"),
				sm.WhereClause("id = ?", 2),
			),
		),
		Into("users", "id", "name"),
		Values(15, "John"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertApply(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES (?, ?), (?, ?)"
	expectedArgs := []any{15, "John", 16, "Mary"}

	query := mysql.Insert(
		Into("users", "id", "name"),
		Apply(func(a mysql.InsertModApply) {
			a.Apply(
				Values(15, "John"),
			)
			a.Apply(
				Values(16, "Mary"),
			)
		}),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertOnDuplicateKey(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES (?, ?), (?, ?) ON DUPLICATE KEY UPDATE name = ?, age = ? RETURNING id"
	expectedArgs := []any{15, "John", 16, "Mary", "Ron", 19}

	query := mysql.Insert(
		Into("users", "id", "name"),
		Values(15, "John"),
		Values(16, "Mary"),
		OnDuplicateKeySet("name", "Ron"),
		OnDuplicateKeySet("age", 19),
		Returning("id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
