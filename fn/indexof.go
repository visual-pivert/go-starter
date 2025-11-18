package fn

import "reflect"

// IndexOf returns the index of the first occurrence of the value in the slice,
// or -1 if the value is not present in the slice.
// Examples:
//
//	IndexOf(1, []int{1, 2, 3}) // 0
func IndexOf[T any](value T, slice []T) int {
	for i, v := range slice {
		if reflect.DeepEqual(v, value) {
			return i
		}
	}
	return -1
}
