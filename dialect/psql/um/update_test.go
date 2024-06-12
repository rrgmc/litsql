package um

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestUpdate(t *testing.T) {
	expectedQuery := "UPDATE users SET name = $1 WHERE id = $2"
	expectedArgs := []any{"John", 15}

	query := psql.Update(
		Table("users"),
		Set("name", "John"),
		WhereClause("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestUpdateBasic(t *testing.T) {
	expectedQuery := "UPDATE ONLY users SET name = $1, created_at = TIME() WHERE id = $2 RETURNING id"
	expectedArgs := []any{"John", 15}

	query := psql.Update(
		Table("users"),
		Only(),
		Set("name", "John"),
		SetString("created_at", "TIME()"),
		WhereClause("id = ?", 15),
		Returning("id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestUpdateFrom(t *testing.T) {
	expectedQuery := "UPDATE users SET address = adr.address, city = ct.city, state = adr.state FROM address AS adr INNER JOIN cities AS ct ON adr.city_id = ct.id WHERE users.address_id = adr.address_id"
	var expectedArgs []any

	query := psql.Update(
		Table("users"),
		From("address AS adr"),
		InnerJoin("cities AS ct").On("adr.city_id = ct.id"),
		SetString("address", "adr.address"),
		SetString("city", "ct.city"),
		SetString("state", "adr.state"),
		WhereClause("users.address_id = adr.address_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestUpdateFromQuery(t *testing.T) {
	expectedQuery := "UPDATE users SET address = adr.address, city = adr.city, state = adr.state FROM (SELECT address, city, state FROM address WHERE id IN ($1, $2, $3)) AS adr WHERE users.address_id = adr.address_id"
	expectedArgs := []any{15, 16, 17}

	query := psql.Update(
		Table("users"),
		FromQuery(
			psql.Select(
				sm.Columns("address", "city", "state"),
				sm.From("address"),
				sm.WhereClause("id IN (?)", expr.In([]any{15, 16, 17})),
			),
		).As("adr"),
		SetString("address", "adr.address"),
		SetString("city", "adr.city"),
		SetString("state", "adr.state"),
		WhereClause("users.address_id = adr.address_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestUpdateWith(t *testing.T) {
	expectedQuery := "WITH city(city_id) AS (SELECT city FROM users WHERE id = $1) UPDATE users SET name = $2, name = city.city_id WHERE id = $3"
	expectedArgs := []any{2, "John", 15}

	query := psql.Update(
		With("city", "city_id").As(
			psql.Select(
				sm.Columns("city"),
				sm.From("users"),
				sm.WhereClause("id = ?", 2),
			),
		),
		Table("users"),
		Set("name", "John"),
		SetString("name", "city.city_id"),
		WhereClause("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestUpdateApply(t *testing.T) {
	expectedQuery := "UPDATE users SET name = $1, age = $2 WHERE id = $3"
	expectedArgs := []any{"John", 51, 15}

	query := psql.Update(
		Table("users"),
		Apply(func(a psql.UpdateModApply) {
			a.Apply(
				Set("name", "John"),
				Set("age", 51),
			)
		}),
		WhereClause("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
