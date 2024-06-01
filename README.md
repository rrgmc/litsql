# litsql - Literal SQL query builder

`litsql` is a Golang string concatenation library **disguised** as an SQL query builder.

Ok, it really is an SQL query builder, but it aims to be an **easier-to-use replacement for raw SQL strings**.

Each `litsql` statement must be directly related to an SQL output, **including whitespace**, which must be obvious to
the user of the library.

The library will do:

 * ensure clause ordering

The library won't do:

 * prevent invalid SQL from being output
 * automatic quoting

## Installation

```shell
go get -u github.com/rrgmc/litsql
```

## Reference

This library is heavily inspired by the excellent [Bob Go SQL Access Toolkit](https://bob.stephenafamo.com/). Its base 
ideas and some of its implementations where used to build this library's interface.

The biggest difference is that `Bob` is not only a query builder, but an ORM, so the query builder part must be
much more complex to be able to tackle multiple jobs. In some parts it is hard to directly relate what SQL will be
generated, and it encourages using Go to code SQL expressions, which this library heavily discourages.

## Author

Rangel Reale (rangelreale@gmail.com)
