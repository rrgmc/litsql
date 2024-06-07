package mysql_test

import (
	"fmt"

	"github.com/rrgmc/litsql/dialect/mysql"
	"github.com/rrgmc/litsql/dialect/mysql/sm"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/sq"
)

func ExampleSelect_literalSimple() {
	// SELECT
	q := mysql.Select(
		// u.id, u.name
		sm.Columns("u.id", "u.name"),
		// , u.created_at, u.updated_at
		sm.Columns("u.created_at", "u.updated_at"),
		// FROM users AS u
		sm.From("users AS u"),
		// WHERE u.age > ?
		sm.WhereC("u.age > ?", 40),
		// WHERE u.city_id = ?
		sm.WhereC("u.city_id = ?", sq.NamedArg("city_id")),
		// AND u.deleted_at IS NOT NULL
		sm.Where("u.deleted_at IS NOT NULL"),
		// ORDER BY u.name ASC, u.age DESC
		sm.OrderBy("u.name ASC", "u.age DESC"),
	)
	qs, args, err := q.Build(
		sq.WithParseArgs(map[string]any{
			"city_id": 66,
		}),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(qs)
	fmt.Println("===")
	fmt.Println(args)

	// Output:
	// SELECT u.id, u.name, u.created_at, u.updated_at
	// FROM users AS u
	// WHERE u.age > ? AND u.city_id = ? AND u.deleted_at IS NOT NULL
	// ORDER BY u.name ASC, u.age DESC
	// ===
	// [40 66]
}

func ExampleSelect_literalSimpleInvalid() {
	// the library won't prevent invalid SQL from being generated, like when using raw string.

	// SELECT
	q := mysql.Select(
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
	q := mysql.Select(
		// orders.id as order_id, orders.date
		sm.Columns("orders.id AS order_id", "orders.date"),
		// u.id AS user_id, u.name as user_name
		sm.Columns("u.id AS user_id", "u.name AS user_name"),
		// FROM orders
		sm.From("orders"),
		// INNER JOIN users AS u ON orders.user_id = u.id
		sm.InnerJoin("users AS u").On("orders.user_id = u.id"),
		// WHERE u.age > ?
		sm.WhereC("u.age ?",
			// example to use either IS NULL or a comparison
			expr.IfElse(true, // some condition
				expr.Clause("> ?", 32),
				expr.String("IS NULL"))),
		// AND u.deleted_at IS NOT NULL
		sm.Where("u.deleted_at IS NOT NULL"),
		// ORDER BY order.date DESC, u.name ASC
		sm.OrderBy("orders.date DESC", "u.name ASC"),
	)
	qs, args, err := q.Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(qs)
	fmt.Println("===")
	fmt.Println(args)

	// Output:
	// SELECT orders.id AS order_id, orders.date, u.id AS user_id, u.name AS user_name
	// FROM orders
	// INNER JOIN users AS u ON orders.user_id = u.id
	// WHERE u.age > ? AND u.deleted_at IS NOT NULL
	// ORDER BY orders.date DESC, u.name ASC
	// ===
	// [32]
}

func ExampleSelect_literalWith() {
	q := mysql.Select(
		// WITH regional_sales AS (
		sm.With("regional_sales").As(
			// SELECT
			mysql.Select(
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
			mysql.Select(
				// region
				sm.Columns("region"),
				// FROM regional_sales
				sm.From("regional_sales"),
				// WHERE total_sales > (SELECT SUM(total_sales)/10 FROM regional_sales)
				sm.WhereC("total_sales > ?",
					mysql.Select(
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
			mysql.Select(
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
