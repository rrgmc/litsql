# litsql - Literal SQL query builder

[![Test Status](https://github.com/rrgmc/litsql/actions/workflows/go.yml/badge.svg)](https://github.com/rrgmc/litsql/actions/workflows/go.yml) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/rrgmc/litsql) [![Go Reference](https://pkg.go.dev/badge/github.com/rrgmc/litsql.svg)](https://pkg.go.dev/github.com/rrgmc/litsql) [![Go Report Card](https://goreportcard.com/badge/github.com/rrgmc/litsql)](https://goreportcard.com/report/github.com/rrgmc/litsql) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/rrgmc/litsql) [![Coverage Status](https://coveralls.io/repos/github/rrgmc/litsql/badge.svg)](https://coveralls.io/github/rrgmc/litsql)

`litsql` is a Golang string concatenation library **disguised as an SQL query builder**.

Ok, it really is an SQL query builder, but it aims to be an **easier-to-use replacement for raw SQL strings**.

Each `litsql` statement **must** be directly related to an SQL output, **including whitespace** (backed by whitespace tests), 
which must be obvious to the user of the library. The output will be exactly the passed values.

```go
func ExampleSelect_literalSimple() {
    // SELECT
    q := psql.Select(
        // u.id, u.name
        sm.Columns("u.id", "u.name"),
        // , u.created_at, u.updated_at
        sm.Columns("u.created_at", "u.updated_at"),
        // FROM users AS u
        sm.From("users AS u"),
        // WHERE u.age > $1
        sm.WhereClause("u.age > ?", 40),
        // WHERE u.city_id = $2
        sm.WhereClause("u.city_id = ?", sq.NamedArg("city_id")),
        // AND u.deleted_at IS NOT NULL
        sm.Where("u.deleted_at IS NOT NULL"),
        // ORDER BY u.name ASC, u.age DESC
        sm.OrderBy("u.name ASC", "u.age DESC"),
    )
    qs, args, err := q.Build(
        sq.WithParseArgs(litsql.MapArgValues{
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
    // WHERE u.age > $1 AND u.city_id = $2 AND u.deleted_at IS NOT NULL
    // ORDER BY u.name ASC, u.age DESC
    // ===
    // [40 66]
}
```

This library will do:

 * ensure clause ordering
 * enforce some kind of code structure
 * be type-safe without using `any` too much
 * guarantee whitespace. **extra whitespace is considered a bug**
 * output correct argument characters for each database dialect

This library won't do:

 * prevent invalid SQL from being output
 * quoting
 * execute queries in databases
 * provide helper expressions to build things like "IsEQ()", "Not(expression)", "LT(value)". These are expected to be written as strings
 * be an ORM (**never**)

The different SQL dialects uses code generation to add/remove things that are dialect specific.

## Installation

```shell
go get -u github.com/rrgmc/litsql
```

## Reference

This library is heavily inspired by the excellent [Bob Go SQL Access Toolkit](https://bob.stephenafamo.com/). Its base
ideas and some of its implementations where used to build this library.

The biggest difference is that `Bob` is not only a query builder, but an ORM, so the query builder part must be
much more complex to be able to tackle multiple jobs. It encourages using Go to code SQL expressions, which this 
library heavily discourages.

## Dialects

* [PostgreSQL](https://pkg.go.dev/github.com/rrgmc/litsql/dialect/psql)
* [MySQL](https://pkg.go.dev/github.com/rrgmc/litsql/dialect/mysql)
* [SQLite](https://pkg.go.dev/github.com/rrgmc/litsql/dialect/sqlite)

## Related

* [libsql-db](https://github.com/rrgmc/litsql-db): wrappers for running `litsql` queries directly in databases.

## Examples

```go
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
        sm.WhereClause("u.age ?",
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
    // WHERE u.age > $1 AND u.deleted_at IS NOT NULL
    // ORDER BY orders.date DESC, u.name ASC
    // ===
    // [32]
}
```


```go
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
                sm.WhereClause("total_sales > ?",
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
        sm.WhereClause("region IN ?",
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
```

## Tasks

#### Dynamic query building

```go
type userFilter struct {
    Name string
}
filter := userFilter{
    Name: "john",
}
query := psql.Select(
    sm.Columns("id", "name"),
    sm.From("users"),
)
if filter.Name != "" {
    query.Apply(
        sm.WhereClause("name = ?", filter.Name),
    )
}
```

#### Select from subselect

```go
query := psql.Select(
    sm.Columns("id", "name", "age"),
    sm.FromQuery(psql.Select(
        sm.Columns("id", "name", "age"),
        sm.From("users"),
        sm.Where("age > 10"),
    )),
)
```

#### WHERE value IN

```go
query := psql.Select(
    sm.Columns("id", "name", "age"),
    sm.From("users"),
    sm.WhereClause("age IN (?)", expr.In([]any{15, 30, 45})),
)
```

#### WHERE value IN using named arguments

```go
query := psql.Select(
    sm.Columns("id", "name", "age"),
    sm.From("users"),
    sm.WhereClause("age IN (?)", expr.In([]any{
        sq.NamedArg("first"),
        sq.NamedArg("second"),
        sq.NamedArg("third"),
    })),
)
qs, args, err := query.Build(
    sq.WithParseArgs(litsql.MapArgValues{
        "first":  15,
        "second": 30,
        "third":  45,
    }),
)
```

#### WHERE value IN subselect

```go
query := psql.Select(
    sm.Columns("id", "name", "age"),
    sm.From("users"),
    sm.WhereClause("region IN ?",
        psql.Select(
            sm.Columns("region"),
            sm.From("top_regions"),
        ),
    ),
)
```

#### Expression function to generate using custom code

```go
query := psql.Select(
    sm.Columns("id", "name", "age"),
    sm.From("users"),
    sm.WhereClause("age > ?",
        expr.Func(func() (litsql.Expression, error) {
            r := rand.Intn(3)
            switch r {
            case 0:
                return expr.Arg(20), nil
            case 1:
                return expr.Arg(30), nil
            default:
                return expr.Arg(50), nil
            }
        }),
    ),
)
```

#### Add clauses in inline callback

```go
query := psql.Select(
    sm.Columns("id", "name", "age"),
    sm.From("users"),
    sm.Apply(func(a psql.SelectModApply) {
        a.Apply(
            sm.Where("age > 10"),
        )
    }),
)
```

#### Use IS NULL or a condition depending on a flag

```go
v := any(32)
query := psql.Select(
    sm.Columns("id", "name", "age"),
    sm.From("users"),
    sm.WhereClause("u.age ?",
        expr.IfElse(v != nil,
            expr.Clause("> ?", 32),
            expr.String("IS NULL"))),
)
```

#### OR expression

```go
query := psql.Select(
    sm.Columns("id", "name", "age"),
    sm.From("users"),
    sm.WhereExpr(
        expr.Or(
            "(age > 10 AND city_id = 12)",
            "(age < 10 AND city_id = 15)",
        ),
    ),
)
```

#### UNION

```go
query := psql.Select(
    sm.Columns("id", "name", "age"),
    sm.From("users"),
    sm.Where("age < 10"),
    sm.Union(psql.Select(
        sm.Columns("id", "name", "age"),
        sm.From("users"),
        sm.Where("age > 50"),
    )),
)
```

#### Full raw query (the query and parameters will be returned as-is)

```go
query := psql.SelectRaw("select * from users where user_id = $1", 55)
```

#### Full raw query (with clause processing)

```go
query := psql.SelectRawExpr(expr.Clause("select * from users where user_id = ?", 55))
```

#### Prepared statements

When using prepared statements, the use of named arguments is required, as it would be impossible to know which
argument maps to each value.

```go
query := psql.Select(
    sm.Columns("film_id", "title", "length"),
    sm.From("film"),
    sm.WhereClause("length > ?", sq.NamedArg("length")),
    sm.LimitArgNamed("limit"),
)

queryStr, args, err := query.Build()
if err != nil {
    return err
}

prepq, err := db.PrepareContext(ctx, queryStr)
if err != nil {
    return err
}

pargs, err := sq.ParseArgs(args, map[string]any{
    "length": 100,
    "limit":  10,
})
if err != nil {
    return err
}

rows, err := prepq.QueryContext(ctx, pargs...)
if err != nil {
    return err
}
defer rows.Close()

for rows.Next() {
    var id, length int
    var title string
    if err := rows.Scan(&id, &title, &length); err != nil {
        return err
    }
    fmt.Println(id, title, length)
}

if rows.Err() != nil {
    return rows.Err()
}
```

## Author

Rangel Reale (rangelreale@gmail.com)
