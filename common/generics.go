package common

// Reduce(fn, initial, []) => initial
func Reduce[T any](fn func(x, y T) T, initial T, values []T) T {
	for _, value := range values {
		initial = fn(initial, value)
	}
	return initial
}

func Map[T any, R any](values []T, fn func(x T) R) []R {
	result := make([]R, len(values))
	for j, v := range values {
		result[j] = fn(v)
	}
	return result
}
