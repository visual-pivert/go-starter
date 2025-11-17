package is

// In checks if the value is in a slice.
// examples:
//
//	is.In(1, []int{1, 2, 3}) // true
//	is.In(4, []int{1, 2, 3}) // false
func In[T comparable](value T, slice []T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
