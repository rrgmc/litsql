package psql_test

import (
	"fmt"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/expr"
)

func ExampleSelect_literalSimple() {
	// SELECT
	q := psql.Select(
		// u.id, u.name
		sm.Columns("u.id", "u.name"),
		// , u.created_at, u.updated_at
		sm.Columns("u.created_at", "u.updated_at"),
		// FROM users As u
		sm.From("users AS u"),
		// WHERE u.age > $1
		sm.WhereC("u.age > ?", 40),
		// AND u.deleted_at IS NOT NULL
		sm.Where("u.deleted_at IS NOT NULL"),
		// ORDER BY u.name ASC, u.age DESC
		sm.OrderBy("u.name ASC", "u.age DESC"),
	)
	qs, args, err := q.Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(qs)
	fmt.Println("===")
	fmt.Println(args)

	// Output:
	// SELECT u.id, u.name, u.created_at, u.updated_at
	// FROM users AS u
	// WHERE u.age > $1 AND u.deleted_at IS NOT NULL
	// ORDER BY u.name ASC, u.age DESC
	// ===
	// [40]
}

func ExampleSelect_literalSimpleInvalid() {
	// the library won't prevent invalid SQL from being generated, like when using raw string.

	// SELECT
	q := psql.Select(
		// u.id, u.name of user
		sm.Columns("u.id", "u.name of user"),
		// FROM users AND NOT users
		sm.From("users AND NOT users"),
		// WHERE CREATE TABLE users
		sm.Where("CREATE TABLE users"),
	)
	qs, _, err := q.Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(qs)

	// Output:
	// SELECT u.id, u.name of user
	// FROM users AND NOT users
	// WHERE CREATE TABLE users
}

func ExampleSelect_literalJoin() {
	// SELECT
	q := psql.Select(
		// orders.id as order_id, orders.date
		sm.Columns("orders.id AS order_id", "orders.date"),
		// u.id AS user_id, u.name as user_name
		sm.Columns("u.id AS user_id", "u.name AS user_name"),
		// FROM orders
		sm.From("orders"),
		// INNER JOIN users AS u ON orders.user_id = u.id
		sm.InnerJoin("users AS u").On("orders.user_id = u.id"),
		// WHERE u.age > $1
		sm.WhereC("u.age ?", expr.F(func() (litsql.Expression, error) {
			// example to use either IS NULL or a comparison
			if true { // some condition
				return expr.C("> ?", 32), nil
			}
			return expr.S("IS NULL"), nil
		})),
		// AND u.deleted_at IS NOT NULL
		sm.Where("u.deleted_at IS NOT NULL"),
		// ORDER BY order.date DESC, u.name ASC
		sm.OrderBy("orders.date DESC", "u.name ASC"),
	)
	qs, _, err := q.Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(qs)

	// Output:
	// SELECT orders.id AS order_id, orders.date, u.id AS user_id, u.name AS user_name
	// FROM orders
	// INNER JOIN users AS u ON orders.user_id = u.id
	// WHERE u.age > $1 AND u.deleted_at IS NOT NULL
	// ORDER BY orders.date DESC, u.name ASC
}

func ExampleSelect_literalWith() {
	q := psql.Select(
		// WITH regional_sales AS (
		sm.With("regional_sales").As(
			// SELECT
			psql.Select(
				// region, SUM(amount) AS total_sales
				sm.Columns("region", "SUM(amount) AS total_sales"),
				// FROM orders
				sm.From("orders"),
				// GROUP BY region
				sm.GroupBy("region"),
			),
		),
		// ), top_regions AS (
		sm.With("top_regions").As(
			// SELECT
			psql.Select(
				// region
				sm.Columns("region"),
				// FROM regional_sales
				sm.From("regional_sales"),
				// WHERE total_sales > (SELECT SUM(total_sales)/10 FROM regional_sales)
				sm.WhereC("total_sales > ?",
					psql.Select(
						sm.Columns("SUM(total_sales)/10"),
						sm.From("regional_sales"),
					),
				),
			),
		),
		// )
		// SELECT
		// region, product, SUM(quantity) AS product_units, SUM(amount) AS product_sales
		sm.Columns("region", "product", "SUM(quantity) AS product_units", "SUM(amount) AS product_sales"),
		// FROM orders
		sm.From("orders"),
		// WHERE region IN (SELECT region FROM top_regions)
		sm.WhereC("region IN ?",
			psql.Select(
				sm.Columns("region"),
				sm.From("top_regions"),
			),
		),
		// GROUP BY region, product
		sm.GroupBy("region", "product"),
	)
	qs, _, err := q.Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(qs)

	// Output:
	// WITH regional_sales AS (
	//   SELECT region, SUM(amount) AS total_sales
	//   FROM orders
	//   GROUP BY region
	// ),
	// top_regions AS (
	//   SELECT region
	//   FROM regional_sales
	//   WHERE total_sales > (
	//     SELECT SUM(total_sales)/10
	//     FROM regional_sales
	//   )
	// )
	// SELECT region, product, SUM(quantity) AS product_units, SUM(amount) AS product_sales
	// FROM orders
	// WHERE region IN (
	//   SELECT region
	//   FROM top_regions
	// )
	// GROUP BY region, product
}
