package fn

func Some[T any](slice []T, fn func(T) bool) bool {
	for _, value := range slice {
		if fn(value) {
			return true
		}
	}
	return false
}
