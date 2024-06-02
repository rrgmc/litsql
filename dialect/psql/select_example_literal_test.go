package psql_test

import (
	"fmt"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
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
		// WHERE u.age > 40
		sm.Where("u.age > 40"),
		// AND u.deleted_at IS NOT NULL
		sm.Where("u.deleted_at IS NOT NULL"),
		// ORDER BY u.name ASC, u.age DESC
		sm.OrderBy("u.name ASC", "u.age DESC"),
	)
	qs, _, err := q.Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(qs)

	// Output:
	// SELECT u.id, u.name, u.created_at, u.updated_at
	// FROM users AS u
	// WHERE u.age > 40 AND u.deleted_at IS NOT NULL
	// ORDER BY u.name ASC, u.age DESC
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
