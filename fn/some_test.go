package fn

import "testing"

func TestSome(t *testing.T) {
	testCases := []struct {
		name     string
		value    []any
		fn       func(any) bool
		expected bool
	}{
		{"contain even number", []any{1, 2, 3}, func(value any) bool { return value.(int)%2 == 0 }, true},
		{"contain odd number", []any{1, 2, 3}, func(value any) bool { return value.(int)%2 == 1 }, true},
		{"not contain 0", []any{1, 2, 3}, func(value any) bool { return value.(int) == 0 }, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := Some(testCase.value, testCase.fn)
			if got != testCase.expected {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
