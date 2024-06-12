package im

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestInsert(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES ($1, $2)"
	expectedArgs := []any{15, "John"}

	query := psql.Insert(
		Into("users", "id", "name"),
		Values(15, "John"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertBasic(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES ($1, $2), ($3, $4) ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name RETURNING id"
	expectedArgs := []any{15, "John", 16, "Mary"}

	query := psql.Insert(
		Into("users", "id", "name"),
		Values(15, "John"),
		Values(16, "Mary"),
		OnConflict("id").DoUpdate(
			ConflictSetString("name", "EXCLUDED.name"),
		),
		Returning("id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertOverriding(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) OVERRIDING SYSTEM VALUE VALUES ($1, $2)"
	expectedArgs := []any{15, "John"}

	query := psql.Insert(
		OverridingSystem(),
		Into("users", "id", "name"),
		Values(15, "John"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertWith(t *testing.T) {
	expectedQuery := "WITH city(city_id) AS (SELECT city FROM users WHERE id = $1) INSERT INTO users (id, name) VALUES ($2, $3)"
	expectedArgs := []any{2, 15, "John"}

	query := psql.Insert(
		With("city", "city_id").As(
			psql.Select(
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
	expectedQuery := "INSERT INTO users (id, name) VALUES ($1, $2), ($3, $4)"
	expectedArgs := []any{15, "John", 16, "Mary"}

	query := psql.Insert(
		Into("users", "id", "name"),
		Apply(func(a psql.InsertModApply) {
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
