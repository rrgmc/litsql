package internal

func ToAnySlice[T any, Ts ~[]T](slice Ts) []any {
	ret := make([]any, len(slice))
	for i, val := range slice {
		ret[i] = val
	}

	return ret
}
