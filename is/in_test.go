package is

import "testing"

func TestIn(t *testing.T) {
	testCases := []struct {
		name     string
		value    any
		slice    []any
		expected bool
	}{
		{"is 'b' in slice", "b", []any{"a", "b", "c"}, true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := In(testCase.value, testCase.slice)
			if got != testCase.expected {
				t.Errorf("got %v, expected %v", got, testCase.expected)
			}
		})
	}
}
