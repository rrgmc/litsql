// Code generated by "litsql-dialectgen"; DO NOT EDIT.
package im

import (
	litsql "github.com/rrgmc/litsql"
	tag "github.com/rrgmc/litsql/dialect/mysql/tag"
	sq "github.com/rrgmc/litsql/sq"
)

type WithChain interface {
	sq.QueryMod[tag.InsertTag]
	Recursive() WithChain
	As(q litsql.Query) WithChain
}
