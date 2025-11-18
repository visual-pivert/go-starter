package fn

// Any returns true if at least one element in the slice satisfies the predicate.
// Examples:
//
//	is.Any([]int{1, 2, 3}, func(v int) bool { return v > 2 }) // true (3 satisfies the predicate)
func Any[T any](slice []T, fn func(T) bool) bool {
	for _, value := range slice {
		if fn(value) {
			return true
		}
	}
	return false
}
