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

func Find[T comparable](f func(T) bool, slice []T) (i int, out T) {
	for i, out = range slice {
		if f(out) {
			return
		}
	}

	var null T
	return -1, null
}
