package fn

func IndexOf[T comparable](value T, slice []T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
