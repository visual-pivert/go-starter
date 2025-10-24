package is

import "testing"

func TestIsFalsy(t *testing.T) {
	testCases := []struct {
		name     string
		value    any
		expected bool
	}{
		{"empty string", "", true},
		{"empty slice", []any{}, true},
		{"not empty slice", []any{"a", "b", "c"}, false},
		{"empty map", map[string]any{}, true},
		{"zero", 0, true},
		{"false", false, true},
		{"nil", nil, true},
		{"empty struct", struct{}{}, true},
		{"not empty struct", struct{ a int }{}, false},
		{"not empty struct", struct{ a int }{1}, false},
		{"zero float", 0.0, true},
		{"zero int", 0, true},
		{"zero int8", int8(0), true},
		{"zero int16", int16(0), true},
		{"zero float32", float32(0), true},
		{"not zero float32", 1, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := Falsy(testCase.value)
			if got != testCase.expected {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
