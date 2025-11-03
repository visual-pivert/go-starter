package is

func In(value any, slice []any) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
