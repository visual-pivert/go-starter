package fn

import "testing"

func TestFn_IndexOf(t *testing.T) {
	testCases := []struct {
		name     string
		value    []any
		target   any
		expected int
	}{
		{"index of 'a'", []any{"a", "b", "c"}, "a", 0},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := IndexOf(testCase.target, testCase.value)
			if got != testCase.expected {
				t.Errorf("got %v, expected %v", got, testCase.expected)
			}
		})
	}
}
