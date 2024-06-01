package mod

import "github.com/rrgmc/litsql/sq"

type InsertConflictUpdateMod[T, UM any] interface {
	sq.Mod[T, UM]
}

func InsertConflictUpdateModFunc[T, UM any](f func(UM)) InsertConflictUpdateMod[T, UM] {
	return &insertConflictUpdateModFunc[T, UM]{f: f}
}

type insertConflictUpdateModFunc[T, UM any] struct {
	sq.ModTagImpl[T]
	f func(UM)
}

func (m insertConflictUpdateModFunc[T, UM]) Apply(v UM) {
	m.f(v)
}
