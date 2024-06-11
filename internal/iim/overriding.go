package iim

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/sq"
)

//litsql:dialects psql
func OverridingSystem[T any]() sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.InsertOverriding{Overriding: "SYSTEM"})
	})
}

//litsql:dialects psql
func OverridingUser[T any]() sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.AddQueryClause(&iclause.InsertOverriding{Overriding: "USER"})
	})
}
