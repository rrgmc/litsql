package dm

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestDelete(t *testing.T) {
	expectedQuery := "DELETE FROM users WHERE id = $1"
	expectedArgs := []any{15}

	query := psql.Delete(
		From("users"),
		WhereClause("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteBasic(t *testing.T) {
	expectedQuery := "DELETE FROM ONLY users WHERE id = $1 RETURNING id"
	expectedArgs := []any{15}

	query := psql.Delete(
		From("users"),
		Only(),
		WhereClause("id = ?", 15),
		Returning("id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteUsing(t *testing.T) {
	expectedQuery := "DELETE FROM users USING address AS adr INNER JOIN cities AS ct ON adr.city_id = ct.id WHERE users.address_id = adr.address_id"
	var expectedArgs []any

	query := psql.Delete(
		From("users"),
		Using("address AS adr"),
		InnerJoin("cities AS ct").On("adr.city_id = ct.id"),
		WhereClause("users.address_id = adr.address_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteUsingQuery(t *testing.T) {
	expectedQuery := "DELETE FROM users USING (SELECT address, city, state FROM address WHERE id IN ($1, $2, $3)) AS adr WHERE users.address_id = adr.address_id"
	expectedArgs := []any{15, 16, 17}

	query := psql.Delete(
		From("users"),
		UsingQuery(
			psql.Select(
				sm.Columns("address", "city", "state"),
				sm.From("address"),
				sm.WhereClause("id IN (?)", expr.In([]any{15, 16, 17})),
			),
		).As("adr"),
		WhereClause("users.address_id = adr.address_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteWith(t *testing.T) {
	expectedQuery := "WITH city(city_id) AS (SELECT city FROM users WHERE id = $1) DELETE FROM users WHERE id = $2"
	expectedArgs := []any{2, 15}

	query := psql.Delete(
		With("city", "city_id").As(
			psql.Select(
				sm.Columns("city"),
				sm.From("users"),
				sm.WhereClause("id = ?", 2),
			),
		),
		From("users"),
		WhereClause("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteApply(t *testing.T) {
	expectedQuery := "DELETE FROM users WHERE id = $1"
	expectedArgs := []any{15}

	query := psql.Delete(
		From("users"),
		Apply(func(a psql.DeleteModApply) {
			a.Apply(
				WhereClause("id = ?", 15),
			)
		}),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
