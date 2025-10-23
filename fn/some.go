package fn

func Some(slice []any, fn func(any) bool) bool {
	for _, value := range slice {
		if fn(value) {
			return true
		}
	}
	return false
}
