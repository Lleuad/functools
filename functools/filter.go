package functools

func Filter[T any](f func(T) bool, slice []T) (out []T) {
	out = make([]T, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			out = append(out, v)
		}
	}

	return
}
