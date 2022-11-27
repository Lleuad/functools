# functools

A collection of functional looking functions I kept rewriting in my Go projects.

## Error Handling
### PassErr
```go
func PassErr[T any](t T, err error) T
```
Ignores `err` and returns `t`.
### LogErr
```go
func LogErr[T any](t T, err error) T
```
Logs `err` with `log.Println` if `err` is non-`nil` and returns `t`.
### FatalErr
```go
func FatalErr[T any](t T, err error) T
```
Logs `err` with `log.Fatalln` if `err` is non-`nil`, otherwise returns `t`.

### PanicErr
```go
func FatalErr[T any](t T, err error) T
```
Logs `err` with `log.Panicln` if `err` is non-`nil`, otherwise returns `t`.

## Reduce
```go
func Reduce[T any](f func(T, T) T, a []T) T
```
Applies `f` cumulatively to the items in `a`, left to right. Similarly to Haskell's `foldl'`

## Map
### Map
```go
func Map[T any, V any](f func(T) V, a []T) []V
```
Maps `f` over the items in `a` one by one.
### MaptoStr
```go
func MaptoStr[T any](a []T) []string
```
Equivalent to `Map(fmt.Sprintf, a)`

## Filter
```go
func Filter[T any](f func(T) bool, a []T) []T
```
Returns a slice of all items `v` in `a` for which `f(v)` returns `true`
