package fn

// Map applies a function to each element of a slice and returns a new slice.
// Examples:
//
//	fn.Map([]int{1, 2, 3}, func(v int) int { return v * 2 }) // [2, 4, 6]
func Map[T any, U any](slice []T, fn func(value T, idx int) U) []U {
	result := make([]U, len(slice))
	for i, value := range slice {
		result[i] = fn(value, i)
	}
	return result
}

// MapReverse applies a function to each element of a slice in reverse order and returns a new slice.
// Examples:
//
//	fn.MapReverse([]int{1, 2, 3}, func(v int) int { return v * 2 }) // [6, 4, 2]
func MapReverse[T any, U any](slice []T, fn func(T, int) U) []U {
	result := make([]U, len(slice))
	for i, value := range slice {
		result[len(slice)-1-i] = fn(value, i)
	}
	return result
}
