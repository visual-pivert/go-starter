package fn

import "reflect"

func IndexOf[T any](value T, slice []T) int {
	for i, v := range slice {
		if reflect.DeepEqual(v, value) {
			return i
		}
	}
	return -1
}
