package iim

import (
	"testing"

	"github.com/rrgmc/litsql/internal/ichain"
	"github.com/rrgmc/litsql/internal/imod"
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
		OnConflict[testutils.TestTag, ichain.InsertConflictUpdate[testutils.TestTag, imod.InsertConflictUpdateModTag]]("id").DoUpdate(
			ConflictSetString[testutils.TestTag, ichain.InsertConflictUpdate[testutils.TestTag, imod.InsertConflictUpdateModTag]]("name", "EXCLUDED.name"),
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
		With[testutils.TestTag, ichain.With[testutils.TestTag]]("city", "city_id").As(
			ism.Select[testutils.TestTag](testutils.NewTestDialect(),
				ism.Columns[testutils.TestTag]("city"),
				ism.From[testutils.TestTag, ichain.From[testutils.TestTag]]("users"),
				ism.WhereClause[testutils.TestTag]("id = ?", 2),
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

func TestInsertOnDuplicateKey(t *testing.T) {
	expectedQuery := "INSERT INTO users (id, name) VALUES ($1, $2), ($3, $4) ON DUPLICATE KEY UPDATE name = $5, age = $6 RETURNING id"
	expectedArgs := []any{15, "John", 16, "Mary", "Ron", 19}

	query := Insert[testutils.TestTag](testutils.NewTestDialect(),
		Into[testutils.TestTag]("users", "id", "name"),
		Values[testutils.TestTag](15, "John"),
		Values[testutils.TestTag](16, "Mary"),
		OnDuplicateKeySet[testutils.TestTag]("name", "Ron"),
		OnDuplicateKeySet[testutils.TestTag]("age", 19),
		Returning[testutils.TestTag]("id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
