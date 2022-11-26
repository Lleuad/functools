package functools

import "log"

func LogErr[T any](t T, err error) T {
	if err != nil {
		log.Println(err)
	}
	return t
}

func FatalErr[T any](t T, err error) T {
	if err != nil {
		log.Fatalln(err)
	}
	return t
}

func PanicErr[T any](t T, err error) T {
	if err != nil {
		log.Panicln(err)
	}
	return t
}
