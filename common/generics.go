package common

// Reduce(fn, initial, []) => initial
func Reduce[T any](fn func(x, y T) T, initial T, values []T) T {
	for _, value := range values {
		initial = fn(initial, value)
	}
	return initial
}
