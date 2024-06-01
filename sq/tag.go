package sq

type ModTag[T any] interface {
	isMod(T)
}

type ModTagImpl[T any] struct {
}

//nolint:unused
func (d ModTagImpl[T]) isMod(T) {}
