package sm

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/internal/testutils"
	"github.com/rrgmc/litsql/sq"
)

func TestSelect(t *testing.T) {
	expectedQuery := "SELECT id, name FROM users WHERE age < 10"
	var expectedArgs []any

	query := psql.Select(
		Columns("id", "name"),
		From("users"),
		Where("age < 10"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectBasic(t *testing.T) {
	expectedQuery := "SELECT DISTINCT u.id, u.name, ua.address FROM users AS u INNER JOIN users_address AS a ON users.id = a.user_id WHERE age < 10 ORDER BY name ASC OFFSET $1 LIMIT $2"
	expectedArgs := []any{10, 100}

	query := psql.Select(
		Distinct(),
		Columns("u.id", "u.name", "ua.address"),
		From("users AS u"),
		InnerJoin("users_address AS a").On("users.id = a.user_id"),
		Where("age < 10"),
		OrderBy("name ASC"),
		Offset(10),
		Limit(100),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectBasicArgs(t *testing.T) {
	expectedQuery := "SELECT DISTINCT u.id, u.name, ua.address FROM users AS u INNER JOIN users_address AS a ON users.id = a.user_id WHERE age < $1 ORDER BY name ASC OFFSET $2 LIMIT $3"
	expectedArgs := []any{10, 10, 100}

	query := psql.Select(
		Distinct(),
		Columns("u.id", "u.name", "ua.address"),
		From("users AS u"),
		InnerJoin("users_address AS a").On("users.id = a.user_id"),
		WhereClause("age < ?", 10),
		OrderBy("name ASC"),
		OffsetArg(10),
		LimitArg(100),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectBasicNamedArgs(t *testing.T) {
	expectedQuery := "SELECT DISTINCT u.id, u.name, ua.address FROM users AS u INNER JOIN users_address AS a ON users.id = a.user_id WHERE age < $1 ORDER BY name ASC OFFSET $2 LIMIT $3"
	expectedArgs := []any{10, 10, 100}

	query := psql.Select(
		Distinct(),
		Columns("u.id", "u.name", "ua.address"),
		From("users AS u"),
		InnerJoin("users_address AS a").On("users.id = a.user_id"),
		WhereClause("age < ?", sq.NamedArg("age")),
		OrderBy("name ASC"),
		OffsetArg(sq.NamedArg("offset")),
		LimitArg(sq.NamedArg("limit")),
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

	query := psql.Select(
		Distinct("id", "name"),
		Columns("id", "name"),
		From("users"),
		Where("age < 10"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectFromQuery(t *testing.T) {
	expectedQuery := "SELECT id, name FROM (SELECT address, city, state FROM user_address ua WHERE ua.user_id = u.id) WHERE age < 10"
	var expectedArgs []any

	query := psql.Select(
		Columns("id", "name"),
		FromQuery(
			psql.Select(
				Columns("address", "city", "state"),
				From("user_address ua"),
				Where("ua.user_id = u.id"),
			),
		),
		Where("age < 10"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectUnion(t *testing.T) {
	expectedQuery := "SELECT id, name FROM users WHERE age < 10 UNION (SELECT id, name FROM users WHERE age > 50)"
	var expectedArgs []any

	query := psql.Select(
		Columns("id", "name"),
		From("users"),
		Where("age < 10"),
		Union(
			psql.Select(
				Columns("id", "name"),
				From("users"),
				Where("age > 50"),
			),
		),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectGroupBy(t *testing.T) {
	expectedQuery := "SELECT country, AVG(age) as age_avg FROM users GROUP BY DISTINCT country HAVING AVG(age) > 10"
	var expectedArgs []any

	query := psql.Select(
		Columns("country", "AVG(age) as age_avg"),
		From("users"),
		GroupBy("country").Distinct(),
		Having("AVG(age) > 10"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectJoin(t *testing.T) {
	expectedQuery := "SELECT u.id, u.name, ai.address FROM users AS u INNER JOIN users_address AS ai ON users.id = ai.user_id LEFT JOIN users_address AS al ON users.id = al.user_id RIGHT JOIN users_address AS ar ON users.id = ar.user_id FULL JOIN users_address AS af ON users.id = af.user_id CROSS JOIN users_address AS ac ON users.id = ac.user_id STRAIGHT_JOIN users_address AS as ON users.id = as.user_id"
	var expectedArgs []any

	query := psql.Select(
		Columns("u.id", "u.name", "ai.address"),
		From("users AS u"),
		InnerJoin("users_address AS ai").On("users.id = ai.user_id"),
		LeftJoin("users_address AS al").On("users.id = al.user_id"),
		RightJoin("users_address AS ar").On("users.id = ar.user_id"),
		FullJoin("users_address AS af").On("users.id = af.user_id"),
		CrossJoin("users_address AS ac").On("users.id = ac.user_id"),
		StraightJoin("users_address AS as").On("users.id = as.user_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectJoinSubSelect(t *testing.T) {
	expectedQuery := "SELECT u.id, u.name, ai.address FROM users AS u INNER JOIN (SELECT address, city, state FROM users_address AS ua WHERE u.id = ua.user_id) AS ua ON users.id = ua.user_id"
	var expectedArgs []any

	query := psql.Select(
		Columns("u.id", "u.name", "ai.address"),
		From("users AS u"),
		InnerJoinExpr(
			psql.Select(
				Columns("address", "city", "state"),
				From("users_address AS ua"),
				Where("u.id = ua.user_id"),
			),
		).As("ua").On("users.id = ua.user_id"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectWith(t *testing.T) {
	expectedQuery := "WITH regional_sales AS (SELECT region, SUM(amount) AS total_sales FROM orders GROUP BY region), top_regions AS (SELECT region FROM regional_sales WHERE total_sales > (SELECT SUM(total_sales)/10 FROM regional_sales)) SELECT region, product, SUM(quantity) AS product_units, SUM(amount) AS product_sales FROM orders WHERE region IN (SELECT region FROM top_regions) GROUP BY region, product"
	var expectedArgs []any

	query := psql.Select(
		With("regional_sales").As(
			psql.Select(
				Columns("region", "SUM(amount) AS total_sales"),
				From("orders"),
				GroupBy("region"),
			),
		),
		With("top_regions").As(
			psql.Select(
				Columns("region"),
				From("regional_sales"),
				WhereClause("total_sales > ?",
					psql.Select(
						Columns("SUM(total_sales)/10"),
						From("regional_sales"),
					),
				),
			),
		),
		Columns("region", "product", "SUM(quantity) AS product_units", "SUM(amount) AS product_sales"),
		From("orders"),
		WhereClause("region IN ?",
			psql.Select(
				Columns("region"),
				From("top_regions"),
			),
		),
		GroupBy("region", "product"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectWindow(t *testing.T) {
	expectedQuery := "SELECT sum(salary) OVER w, avg(salary) OVER w FROM empsalary WINDOW w AS (PARTITION BY depname ORDER BY salary DESC)"
	var expectedArgs []any

	query := psql.Select(
		Columns("sum(salary) OVER w", "avg(salary) OVER w"),
		From("empsalary"),
		Window("w").
			PartitionBy("depname").
			OrderBy("salary DESC"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}

func TestSelectApply(t *testing.T) {
	expectedQuery := "SELECT id, name FROM users WHERE age < 10 ORDER BY name"
	var expectedArgs []any

	query := psql.Select(
		Columns("id", "name"),
		Apply(func(a psql.SelectModApply) {
			a.Apply(
				Where("age < 10"),
				OrderBy("name"),
			)
		}),
		From("users"),
	)

	testutils.TestQuery(t, query, expectedQuery, expectedArgs...)
}
