package is

import "testing"

func TestIs_Truthy(t *testing.T) {
	testCases := []struct {
		name     string
		value    any
		expected bool
	}{

		{"empty string", "", false},
		{"empty slice", []any{}, false},
		{"not empty slice", []any{"a", "b", "c"}, true},
		{"empty map", map[string]any{}, false},
		{"zero", 0, false},
		{"false", false, false},
		{"nil", nil, false},
		{"empty struct", struct{}{}, false},
		{"not empty struct", struct{ a int }{}, true},
		{"not empty struct", struct{ a int }{1}, true},
		{"zero float", 0.0, false},
		{"zero int", 0, false},
		{"zero int8", int8(0), false},
		{"zero int16", int16(0), false},
		{"zero float32", float32(0), false},
		{"not zero float32", float32(1), true},
		{"greater than 0", 20, true},
		{"greater than 1", 21, true},
		{"str length greater or equal to 1", "qwe", true},
		{"str length greater or equal to 2", "qw", true},
		{"str length greater or equal to 3", "q", true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := Truthy(testCase.value)
			if got != testCase.expected {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
