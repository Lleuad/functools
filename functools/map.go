package functools

import "fmt"

func Map[T any, V any](f func(T) V, slice []T) (out []V) {
	out = make([]V, len(slice))
	for i, v := range slice {
		out[i] = f(v)
	}

	return
}

func MaptoStr[T any](slice []T) (out []string) {
	out = make([]string, len(slice))
	for i, v := range slice {
		out[i] = fmt.Sprint(v)
	}

	return
}
