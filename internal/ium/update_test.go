package ium

import (
	"testing"

	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/ism"
	"github.com/rrgmc/litsql/internal/testutils"
	"github.com/rrgmc/litsql/sq"
)

func TestUpdate(t *testing.T) {
	expectedQuery := "UPDATE users SET name = $1 WHERE id = $2"
	expectedArgs := []any{"John", 15}

	query := Update[testutils.TestTag](testutils.NewTestDialect(),
		Table[testutils.TestTag]("users"),
		Set[testutils.TestTag]("name", "John"),
		WhereC[testutils.TestTag]("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestUpdateBasic(t *testing.T) {
	expectedQuery := "UPDATE ONLY users SET name = $1, created_at = TIME() WHERE id = $2 RETURNING id"
	expectedArgs := []any{"John", 15}

	query := Update[testutils.TestTag](testutils.NewTestDialect(),
		Table[testutils.TestTag]("users"),
		Only[testutils.TestTag](true),
		Set[testutils.TestTag]("name", "John"),
		SetS[testutils.TestTag]("created_at", "TIME()"),
		WhereC[testutils.TestTag]("id = ?", 15),
		Returning[testutils.TestTag]("id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestUpdateFrom(t *testing.T) {
	expectedQuery := "UPDATE users SET address = adr.address, city = ct.city, state = adr.state FROM address AS adr INNER JOIN cities AS ct ON adr.city_id = ct.id WHERE users.address_id = adr.address_id"
	var expectedArgs []any

	d := testutils.NewTestDialect()

	query := Update[testutils.TestTag](d,
		Table[testutils.TestTag]("users"),
		From[testutils.TestTag]("address AS adr"),
		InnerJoin[testutils.TestTag]("cities AS ct").On("adr.city_id = ct.id"),
		SetS[testutils.TestTag]("address", "adr.address"),
		SetS[testutils.TestTag]("city", "ct.city"),
		SetS[testutils.TestTag]("state", "adr.state"),
		WhereC[testutils.TestTag]("users.address_id = adr.address_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestUpdateFromQuery(t *testing.T) {
	expectedQuery := "UPDATE users SET address = adr.address, city = adr.city, state = adr.state FROM (SELECT address, city, state FROM address WHERE id IN ($1, $2, $3)) AS adr WHERE users.address_id = adr.address_id"
	expectedArgs := []any{15, 16, 17}

	d := testutils.NewTestDialect()

	query := Update[testutils.TestTag](d,
		Table[testutils.TestTag]("users"),
		FromQ[testutils.TestTag, testutils.TestTag](
			ism.Select[testutils.TestTag](d,
				ism.Columns[testutils.TestTag]("address", "city", "state"),
				ism.From[testutils.TestTag]("address"),
				ism.WhereClause[testutils.TestTag]("id IN (?)", expr.In([]any{15, 16, 17})),
			),
		).As("adr"),
		SetS[testutils.TestTag]("address", "adr.address"),
		SetS[testutils.TestTag]("city", "adr.city"),
		SetS[testutils.TestTag]("state", "adr.state"),
		WhereC[testutils.TestTag]("users.address_id = adr.address_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestUpdateWith(t *testing.T) {
	expectedQuery := "WITH city(city_id) AS (SELECT city FROM users WHERE id = $1) UPDATE users SET name = $2, name = city.city_id WHERE id = $3"
	expectedArgs := []any{2, "John", 15}

	query := Update[testutils.TestTag](testutils.NewTestDialect(),
		With[testutils.TestTag]("city", "city_id").As(
			ism.Select[testutils.TestTag](testutils.NewTestDialect(),
				ism.Columns[testutils.TestTag]("city"),
				ism.From[testutils.TestTag]("users"),
				ism.WhereClause[testutils.TestTag]("id = ?", 2),
			),
		),
		Table[testutils.TestTag]("users"),
		Set[testutils.TestTag]("name", "John"),
		SetS[testutils.TestTag]("name", "city.city_id"),
		WhereC[testutils.TestTag]("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestUpdateApply(t *testing.T) {
	expectedQuery := "UPDATE users SET name = $1, age = $2 WHERE id = $3"
	expectedArgs := []any{"John", 51, 15}

	query := Update[testutils.TestTag](testutils.NewTestDialect(),
		Table[testutils.TestTag]("users"),
		Apply[testutils.TestTag](func(a sq.QueryModApply[testutils.TestTag]) {
			a.Apply(
				Set[testutils.TestTag]("name", "John"),
				Set[testutils.TestTag]("age", 51),
			)
		}),
		WhereC[testutils.TestTag]("id = ?", 15),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
