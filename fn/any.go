package fn

func Any(slice []any, fn func(any) bool) bool {
	for _, value := range slice {
		if !fn(value) {
			return false
		}
	}
	return true
}
