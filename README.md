# litsql - Literal SQL query builder

`litsql` is a Golang string concatenation library **disguised** as an SQL query builder.

Ok, it really is an SQL query builder, but it aims to be an **easier-to-use replacement for raw SQL strings**.

Each `litsql` statement must be directly related to an SQL output, **including whitespace**, which must be obvious to
the user of the library. The output will be exactly the passed values, so the library won't prevent invalid SQL from
being generated.

The library tests **includes testing for exact string and whitespace output** to ensure this.

```go
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
```

The library will do:

 * ensure clause ordering

The library won't do:

 * prevent invalid SQL from being output
 * automatic quoting

## Installation

```shell
go get -u github.com/rrgmc/litsql
```

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
        sm.WhereC("u.age ?",
            // example to use either IS NULL or a comparison
            expr.ExprIfElse(true, // some condition
                expr.C("> ?", 32),
                expr.S("IS NULL"))),
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
```

## Reference

This library is heavily inspired by the excellent [Bob Go SQL Access Toolkit](https://bob.stephenafamo.com/). Its base 
ideas and some of its implementations where used to build this library's interface.

The biggest difference is that `Bob` is not only a query builder, but an ORM, so the query builder part must be
much more complex to be able to tackle multiple jobs. In some parts it is hard to directly relate what SQL will be
generated, and it encourages using Go to code SQL expressions, which this library heavily discourages.

## Author

Rangel Reale (rangelreale@gmail.com)
