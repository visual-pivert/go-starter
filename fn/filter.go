package fn

func Filter[T any](slice []T, fn func(T) bool) []T {
	var out []T
	for _, value := range slice {
		if fn(value) {
			out = append(out, value)
		}
	}
	return out
}

func FilterI[T any](slice []T, fn func(T) bool) []int {
	var out []int
	for i, value := range slice {
		if fn(value) {
			out = append(out, i)
		}
	}
	return out
}
