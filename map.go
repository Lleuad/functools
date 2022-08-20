package functools

import "fmt"

func Map[T any, V any](f func(T) V, slice []T) (out []V) {
	for _, v := range slice {
		out = append(out, f(v))
	}

	return
}

func MaptoStr[T any](slice []T) (out []string) {
	for _, v := range slice {
		out = append(out, fmt.Sprint(v))
	}

	return
}
