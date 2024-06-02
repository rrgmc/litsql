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
