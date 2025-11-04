package fn

func Map[T any, U any](slice []T, fn func(value T, idx int) U) []U {
	result := make([]U, len(slice))
	for i, value := range slice {
		result[i] = fn(value, i)
	}
	return result
}

func MapReverse[T any, U any](slice []T, fn func(T, int) U) []U {
	result := make([]U, len(slice))
	for i, value := range slice {
		result[len(slice)-1-i] = fn(value, i)
	}
	return result
}
