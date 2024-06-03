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
		WhereC[testutils.TestTag]("age < ?", sq.NamedArg("age")),
		OrderBy[testutils.TestTag]("name ASC"),
		OffsetA[testutils.TestTag](sq.NamedArg("offset")),
		LimitA[testutils.TestTag](sq.NamedArg("limit")),
	)

	testutils.TestQueryParseArgs(t, query, expectedQuery, map[string]any{
		"age":    10,
		"offset": 10,
		"limit":  100,
	}, expectedArgs...)
}

func TestSelectDistinctOn(t *testing.T) {
	expectedQuery := "SELECT DISTINCT ON (id, name) id, name FROM users WHERE age < 10"
	var expectedArgs []any

	query := Select[testutils.TestTag](testutils.NewTestDialect(),
		Distinct[testutils.TestTag]("id", "name"),
		Columns[testutils.TestTag]("id", "name"),
		From[testutils.TestTag]("users"),
		Where[testutils.TestTag]("age < 10"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectFromQuery(t *testing.T) {
	expectedQuery := "SELECT id, name FROM (SELECT address, city, state FROM user_address ua WHERE ua.user_id = u.id) WHERE age < 10"
	var expectedArgs []any

	d := testutils.NewTestDialect()

	query := Select[testutils.TestTag](d,
		Columns[testutils.TestTag]("id", "name"),
		FromQ[testutils.TestTag, testutils.TestTag](
			Select[testutils.TestTag](d,
				Columns[testutils.TestTag]("address", "city", "state"),
				From[testutils.TestTag]("user_address ua"),
				Where[testutils.TestTag]("ua.user_id = u.id"),
			),
		),
		Where[testutils.TestTag]("age < 10"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectUnion(t *testing.T) {
	expectedQuery := "SELECT id, name FROM users WHERE age < 10 UNION (SELECT id, name FROM users WHERE age > 50)"
	var expectedArgs []any

	d := testutils.NewTestDialect()

	query := Select[testutils.TestTag](d,
		Columns[testutils.TestTag]("id", "name"),
		From[testutils.TestTag]("users"),
		Where[testutils.TestTag]("age < 10"),
		Union[testutils.TestTag](
			Select[testutils.TestTag](d,
				Columns[testutils.TestTag]("id", "name"),
				From[testutils.TestTag]("users"),
				Where[testutils.TestTag]("age > 50"),
			),
		),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectGroupBy(t *testing.T) {
	expectedQuery := "SELECT country, AVG(age) as age_avg FROM users GROUP BY country HAVING AVG(age) > 10"
	var expectedArgs []any

	query := Select[testutils.TestTag](testutils.NewTestDialect(),
		Columns[testutils.TestTag]("country", "AVG(age) as age_avg"),
		From[testutils.TestTag]("users"),
		GroupBy[testutils.TestTag]("country"),
		Having[testutils.TestTag]("AVG(age) > 10"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectJoin(t *testing.T) {
	expectedQuery := "SELECT u.id, u.name, ai.address FROM users AS u INNER JOIN users_address AS ai ON users.id = ai.user_id LEFT JOIN users_address AS al ON users.id = al.user_id RIGHT JOIN users_address AS ar ON users.id = ar.user_id FULL JOIN users_address AS af ON users.id = af.user_id CROSS JOIN users_address AS ac ON users.id = ac.user_id STRAIGHT_JOIN users_address AS as ON users.id = as.user_id"
	var expectedArgs []any

	query := Select[testutils.TestTag](testutils.NewTestDialect(),
		Columns[testutils.TestTag]("u.id", "u.name", "ai.address"),
		From[testutils.TestTag]("users AS u"),
		InnerJoin[testutils.TestTag]("users_address AS ai").On("users.id = ai.user_id"),
		LeftJoin[testutils.TestTag]("users_address AS al").On("users.id = al.user_id"),
		RightJoin[testutils.TestTag]("users_address AS ar").On("users.id = ar.user_id"),
		FullJoin[testutils.TestTag]("users_address AS af").On("users.id = af.user_id"),
		CrossJoin[testutils.TestTag]("users_address AS ac").On("users.id = ac.user_id"),
		StraightJoin[testutils.TestTag]("users_address AS as").On("users.id = as.user_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectJoinSubSelect(t *testing.T) {
	expectedQuery := "SELECT u.id, u.name, ai.address FROM users AS u INNER JOIN (SELECT address, city, state FROM users_address AS ua WHERE u.id = ua.user_id) AS ua ON users.id = ua.user_id"
	var expectedArgs []any

	d := testutils.NewTestDialect()

	query := Select[testutils.TestTag](d,
		Columns[testutils.TestTag]("u.id", "u.name", "ai.address"),
		From[testutils.TestTag]("users AS u"),
		InnerJoinE[testutils.TestTag](
			Select[testutils.TestTag](d,
				Columns[testutils.TestTag]("address", "city", "state"),
				From[testutils.TestTag]("users_address AS ua"),
				Where[testutils.TestTag]("u.id = ua.user_id"),
			),
		).As("ua").On("users.id = ua.user_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectWith(t *testing.T) {
	expectedQuery := "WITH regional_sales AS (SELECT region, SUM(amount) AS total_sales FROM orders GROUP BY region), top_regions AS (SELECT region FROM regional_sales WHERE total_sales > (SELECT SUM(total_sales)/10 FROM regional_sales)) SELECT region, product, SUM(quantity) AS product_units, SUM(amount) AS product_sales FROM orders WHERE region IN (SELECT region FROM top_regions) GROUP BY region, product"
	var expectedArgs []any

	d := testutils.NewTestDialect()

	query := Select[testutils.TestTag](d,
		With[testutils.TestTag]("regional_sales").As(
			Select[testutils.TestTag](d,
				Columns[testutils.TestTag]("region", "SUM(amount) AS total_sales"),
				From[testutils.TestTag]("orders"),
				GroupBy[testutils.TestTag]("region"),
			),
		),
		With[testutils.TestTag]("top_regions").As(
			Select[testutils.TestTag](d,
				Columns[testutils.TestTag]("region"),
				From[testutils.TestTag]("regional_sales"),
				WhereC[testutils.TestTag]("total_sales > ?",
					Select[testutils.TestTag](d,
						Columns[testutils.TestTag]("SUM(total_sales)/10"),
						From[testutils.TestTag]("regional_sales"),
					),
				),
			),
		),
		Columns[testutils.TestTag]("region", "product", "SUM(quantity) AS product_units", "SUM(amount) AS product_sales"),
		From[testutils.TestTag]("orders"),
		WhereC[testutils.TestTag]("region IN ?",
			Select[testutils.TestTag](d,
				Columns[testutils.TestTag]("region"),
				From[testutils.TestTag]("top_regions"),
			),
		),
		GroupBy[testutils.TestTag]("region", "product"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectWindow(t *testing.T) {
	expectedQuery := "SELECT sum(salary) OVER w, avg(salary) OVER w FROM empsalary WINDOW w AS (PARTITION BY depname ORDER BY salary DESC)"
	var expectedArgs []any

	query := Select[testutils.TestTag](testutils.NewTestDialect(),
		Columns[testutils.TestTag]("sum(salary) OVER w", "avg(salary) OVER w"),
		From[testutils.TestTag]("empsalary"),
		Window[testutils.TestTag]("w").
			PartitionBy("depname").
			OrderBy("salary DESC"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectApply(t *testing.T) {
	expectedQuery := "SELECT id, name FROM users WHERE age < 10 ORDER BY name"
	var expectedArgs []any

	query := Select[testutils.TestTag](testutils.NewTestDialect(),
		Columns[testutils.TestTag]("id", "name"),
		Apply[testutils.TestTag](func(a sq.QueryModApply[testutils.TestTag]) {
			a.Apply(
				Where[testutils.TestTag]("age < 10"),
				OrderBy[testutils.TestTag]("name"),
			)
		}),
		From[testutils.TestTag]("users"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
