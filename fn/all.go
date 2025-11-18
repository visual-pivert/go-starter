package fn

// All returns true if all elements in the slice satisfy the predicate.
// Examples:
//
//	fn.All([]int{1, 2, 3}, func(v int) bool { return v > 0 }) // true
func All[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}
