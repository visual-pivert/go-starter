package fn

// Reduce returns a new slice with cumulative results.
// Examples:
//
//	fn.Reduce([]int{1,2,3}, 0, func(cum int, value int, index int) int {
//		return cum + value
//	}) // [1, 3, 6] (v[len(v)-1] = sum of 1,2,3)
func Reduce[T any](slice []T, initialValue T, fn func(cum T, value T, index int) T) []T {
	out := make([]T, len(slice))
	out[0] = fn(initialValue, slice[0], 0)
	for i := 1; i < len(slice); i++ {
		out[i] = fn(out[i-1], slice[i], i)
	}
	return out
}
