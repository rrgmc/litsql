package ism

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/iclause"
	"github.com/rrgmc/litsql/internal/isq"
	"github.com/rrgmc/litsql/sq"
)

func Union[T any](q isq.Query[T]) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Combine{
			Strategy: iclause.Union,
			Query:    q,
			All:      false,
		})
	})
}

func UnionAll[T any](q isq.Query[T]) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Combine{
			Strategy: iclause.Union,
			Query:    q,
			All:      true,
		})
	})
}

func Intersect[T any](q isq.Query[T]) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Combine{
			Strategy: iclause.Intersect,
			Query:    q,
			All:      false,
		})
	})
}

func IntersectAll[T any](q isq.Query[T]) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Combine{
			Strategy: iclause.Intersect,
			Query:    q,
			All:      true,
		})
	})
}

func Except[T any](q isq.Query[T]) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Combine{
			Strategy: iclause.Except,
			Query:    q,
			All:      false,
		})
	})
}

func ExceptAll[T any](q isq.Query[T]) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		a.Add(&iclause.Combine{
			Strategy: iclause.Except,
			Query:    q,
			All:      true,
		})
	})
}
