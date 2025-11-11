package fn

func Reduce[T any](slice []T, initialValue T, fn func(cum T, value T, index int) T) []T {
	out := make([]T, len(slice))
	out[0] = fn(initialValue, slice[0], 0)
	for i := 1; i < len(slice); i++ {
		out[i] = fn(out[i-1], slice[i], i)
	}
	return out
}
