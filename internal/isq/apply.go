package isq

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

type funcApply[T any] struct {
	mods []sq.QueryMod[T]
}

func (f *funcApply[T]) Apply(mod ...sq.QueryMod[T]) {
	f.mods = append(f.mods, mod...)
}

func Apply[T any](f func(a sq.QueryModApply[T])) sq.QueryMod[T] {
	return sq.QueryModFunc[T](func(a litsql.QueryBuilder) {
		fi := &funcApply[T]{}
		f(fi)
		for _, m := range fi.mods {
			m.Apply(a)
		}
	})
}
