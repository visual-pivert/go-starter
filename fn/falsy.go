package fn

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
	return isZero(value)
}

func isZero(value any) bool {
	switch v := value.(type) {
	case int:
		return v == 0
	case int8:
		return v == 0
	case int16:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0
	case uint:
		return v == 0
	case byte:
		return v == 0
	case float32:
		return v == 0
	case float64:
		return v == 0
	case complex64:
		return v == 0
	case complex128:
		return v == 0
	default:
		return false
	}
}
