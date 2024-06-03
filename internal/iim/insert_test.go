package iim

import (
	"testing"

	"github.com/rrgmc/litsql/internal/ism"
	"github.com/rrgmc/litsql/internal/testutils"
	"github.com/rrgmc/litsql/sq"
)

func TestInsert(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES ($1, $2)"
	expectedArgs := []any{15, "John"}

	query := Insert[testutils.TestTag](testutils.NewTestDialect(),
		Into[testutils.TestTag]("users", "id", "name"),
		Values[testutils.TestTag](15, "John"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertBasic(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES ($1, $2), ($3, $4) ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name RETURNING id"
	expectedArgs := []any{15, "John", 16, "Mary"}

	query := Insert[testutils.TestTag](testutils.NewTestDialect(),
		Into[testutils.TestTag]("users", "id", "name"),
		Values[testutils.TestTag](15, "John"),
		Values[testutils.TestTag](16, "Mary"),
		OnConflict[testutils.TestTag]("id").DoUpdate(
			ConflictSetS[testutils.TestTag]("name", "EXCLUDED.name"),
		),
		Returning[testutils.TestTag]("id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertOverriding(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) OVERRIDING SYSTEM VALUE VALUES ($1, $2)"
	expectedArgs := []any{15, "John"}

	query := Insert[testutils.TestTag](testutils.NewTestDialect(),
		OverridingSystem[testutils.TestTag](),
		Into[testutils.TestTag]("users", "id", "name"),
		Values[testutils.TestTag](15, "John"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertWith(t *testing.T) {
	expectedQuery := "WITH city(city_id) AS (SELECT city FROM users WHERE id = $1) INSERT INTO users (id, name) VALUES ($2, $3)"
	expectedArgs := []any{2, 15, "John"}

	query := Insert[testutils.TestTag](testutils.NewTestDialect(),
		With[testutils.TestTag]("city", "city_id").As(
			ism.Select[testutils.TestTag](testutils.NewTestDialect(),
				ism.Columns[testutils.TestTag]("city"),
				ism.From[testutils.TestTag]("users"),
				ism.WhereC[testutils.TestTag]("id = ?", 2),
			),
		),
		Into[testutils.TestTag]("users", "id", "name"),
		Values[testutils.TestTag](15, "John"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestInsertApply(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES ($1, $2), ($3, $4)"
	expectedArgs := []any{15, "John", 16, "Mary"}

	query := Insert[testutils.TestTag](testutils.NewTestDialect(),
		Into[testutils.TestTag]("users", "id", "name"),
		Apply[testutils.TestTag](func(a sq.QueryModApply[testutils.TestTag]) {
			a.Apply(
				Values[testutils.TestTag](15, "John"),
			)
			a.Apply(
				Values[testutils.TestTag](16, "Mary"),
			)
		}),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
