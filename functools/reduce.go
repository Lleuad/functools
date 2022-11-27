package functools

func Reduce[T any](f func(T, T) T, slice []T) (out T) {
	out = slice[0]
	for _, v := range slice[1:] {
		out = f(out, v)
	}

	return
}
