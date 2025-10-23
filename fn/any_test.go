package fn

import "testing"

func TestAny(t *testing.T) {
	testCases := []struct {
		name     string
		value    []any
		fn       func(any) bool
		expected any
	}{
		{"greater than 0", []any{1, 2, 3}, func(a any) bool { return a.(int) > 0 }, true},
		{"greater than 1", []any{1, 2, 3}, func(a any) bool { return a.(int) > 1 }, false},
		{"str length greater or equal to 1", []any{"qwe", "qw", "q"}, func(a any) bool { return len(a.(string)) >= 1 }, true},
		{"str length greater or equal to 2", []any{"qwe", "qw", "q"}, func(a any) bool { return len(a.(string)) >= 2 }, false},
		{"str length greater or equal to 3", []any{"qwe", "qw", "q"}, func(a any) bool { return len(a.(string)) >= 3 }, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := Any(testCase.value, testCase.fn)
			if got != testCase.expected {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
