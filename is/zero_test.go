package is

import "testing"

func TestIs_Zero(t *testing.T) {
	testCases := []struct {
		name     string
		value    any
		expected bool
	}{
		{"zero", 0, true},
		{"zero float", 0.0, true},
		{"not zero", 1, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := Zero(testCase.value)
			if got != testCase.expected {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
