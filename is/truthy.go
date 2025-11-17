package is

// Truthy returns true if the value is truthy.
// 1, "abc", true, map[string]int{"a": 1} are truthy values.
// examples:
//
//	is.Truthy(1) // true
func Truthy(value any) bool {
	return !Falsy(value)
}
