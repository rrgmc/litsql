package sq

// ModTag is used just for interface tagging, so implementations for each dialect aren't accept on others.
type ModTag[T any] interface {
	isMod(T)
}

type ModTagImpl[T any] struct {
}

//nolint:unused
func (d ModTagImpl[T]) isMod(T) {}
