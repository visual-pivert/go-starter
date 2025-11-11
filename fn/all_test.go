package fn

import "testing"

func TestAll(t *testing.T) {
	testCases := []struct {
		name     string
		value    []any
		expected bool
	}{
		{"all 2", []any{2, 2, 2}, true},
		{"not all 2", []any{2, 2, 1}, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := All(testCase.value, func(value any) bool {
				return value.(int) == 2
			})
			if got != testCase.expected {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
