package is

// Falsy returns true if the value is falsy.
// 0, "", false, nil, []int{}, map[string]int{} are falsy values.
// examples:
//
//	is.Falsy(0) // true
//	is.Falsy("") // true
func Falsy(value any) bool {
	if value == nil {
		return true
	}

	// Check collection types
	if m, ok := value.(map[string]any); ok && len(m) == 0 {
		return true
	}
	if arr, ok := value.([]any); ok && len(arr) == 0 {
		return true
	}
	if str, ok := value.(string); ok && len(str) == 0 {
		return true
	}

	// Check struct types
	if _, ok := value.(struct{}); ok {
		return true
	}

	// Check boolean
	if b, ok := value.(bool); ok && !b {
		return true
	}

	// Check numeric types
	return Zero(value)
}
