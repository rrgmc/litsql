package psql_test

import (
	"fmt"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/sm"
	"github.com/rrgmc/litsql/sq"
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
	qs, _, err := sq.Build(q)
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
