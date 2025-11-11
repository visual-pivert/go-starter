package fn

func Reverse[T any](slice []T) []T {
	out := make([]T, 0, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		out = append(out, slice[i])
	}
	return out
}
