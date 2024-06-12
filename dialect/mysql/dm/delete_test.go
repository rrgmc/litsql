package dm

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/mysql"
	"github.com/rrgmc/litsql/dialect/mysql/sm"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestDelete(t *testing.T) {
	expectedQuery := "DELETE FROM users WHERE id = $1"
	expectedArgs := []any{15}

	query := mysql.Delete(
		From("users"),
		WhereClause("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteBasic(t *testing.T) {
	expectedQuery := "DELETE FROM ONLY users WHERE id = $1 RETURNING id"
	expectedArgs := []any{15}

	query := mysql.Delete(
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

	query := mysql.Delete(
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

	query := mysql.Delete(
		From("users"),
		UsingQuery(
			mysql.Select(
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

	query := mysql.Delete(
		With("city", "city_id").As(
			mysql.Select(
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

	query := mysql.Delete(
		From("users"),
		Apply(func(a mysql.DeleteModApply) {
			a.Apply(
				WhereClause("id = ?", 15),
			)
		}),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
