package sq

type QueryModApply[T any] interface {
	Apply(...QueryMod[T])
}
