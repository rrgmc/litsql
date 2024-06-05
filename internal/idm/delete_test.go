package idm

import (
	"testing"

	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ism"
	"github.com/rrgmc/litsql/internal/testutils"
	"github.com/rrgmc/litsql/sq"
)

func TestDelete(t *testing.T) {
	expectedQuery := "DELETE FROM users WHERE id = $1"
	expectedArgs := []any{15}

	query := Delete[testutils.TestTag](testutils.NewTestDialect(),
		From[testutils.TestTag]("users"),
		WhereC[testutils.TestTag]("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteBasic(t *testing.T) {
	expectedQuery := "DELETE FROM ONLY users WHERE id = $1 RETURNING id"
	expectedArgs := []any{15}

	query := Delete[testutils.TestTag](testutils.NewTestDialect(),
		From[testutils.TestTag]("users"),
		Only[testutils.TestTag](true),
		WhereC[testutils.TestTag]("id = ?", 15),
		Returning[testutils.TestTag]("id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteUsing(t *testing.T) {
	expectedQuery := "DELETE FROM users USING address AS adr INNER JOIN cities AS ct ON adr.city_id = ct.id WHERE users.address_id = adr.address_id"
	var expectedArgs []any

	d := testutils.NewTestDialect()

	query := Delete[testutils.TestTag](d,
		From[testutils.TestTag]("users"),
		Using[testutils.TestTag]("address AS adr"),
		InnerJoin[testutils.TestTag]("cities AS ct").On("adr.city_id = ct.id"),
		WhereC[testutils.TestTag]("users.address_id = adr.address_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteUsingQuery(t *testing.T) {
	expectedQuery := "DELETE FROM users USING (SELECT address, city, state FROM address WHERE id IN ($1, $2, $3)) AS adr WHERE users.address_id = adr.address_id"
	expectedArgs := []any{15, 16, 17}

	d := testutils.NewTestDialect()

	query := Delete[testutils.TestTag](d,
		From[testutils.TestTag]("users"),
		UsingQ[testutils.TestTag, testutils.TestTag](
			ism.Select[testutils.TestTag](d,
				ism.Columns[testutils.TestTag]("address", "city", "state"),
				ism.From[testutils.TestTag]("address"),
				ism.WhereC[testutils.TestTag]("id IN (?)", expr.In([]any{15, 16, 17})),
			),
		).As("adr"),
		WhereC[testutils.TestTag]("users.address_id = adr.address_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteWith(t *testing.T) {
	expectedQuery := "WITH city(city_id) AS (SELECT city FROM users WHERE id = $1) DELETE FROM users WHERE id = $2"
	expectedArgs := []any{2, 15}

	query := Delete[testutils.TestTag](testutils.NewTestDialect(),
		With[testutils.TestTag]("city", "city_id").As(
			ism.Select[testutils.TestTag](testutils.NewTestDialect(),
				ism.Columns[testutils.TestTag]("city"),
				ism.From[testutils.TestTag]("users"),
				ism.WhereC[testutils.TestTag]("id = ?", 2),
			),
		),
		From[testutils.TestTag]("users"),
		WhereC[testutils.TestTag]("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestDeleteApply(t *testing.T) {
	expectedQuery := "DELETE FROM users WHERE id = $1"
	expectedArgs := []any{15}

	query := Delete[testutils.TestTag](testutils.NewTestDialect(),
		From[testutils.TestTag]("users"),
		Apply[testutils.TestTag](func(a sq.QueryModApply[testutils.TestTag]) {
			a.Apply(
				WhereC[testutils.TestTag]("id = ?", 15),
			)
		}),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
