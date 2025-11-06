package series

import "testing"

func TestSeries_Agg(t *testing.T) {
	testCases := []struct {
		name         string
		initialValue any
		value        []any
		fn           func(any, any, int) any
		expected     any
	}{
		{"sum with aggregate", 0, []any{1, 2, 3}, func(a, b any, _ int) any { return a.(int) + b.(int) }, 6},
		{"multiply with aggregate", 1, []any{1, 2, 3}, func(a, b any, _ int) any { return a.(int) * b.(int) }, 6},
		{"multiply with his index and sum with aggregate", 0, []any{1, 2, 3}, func(a, b any, index int) any { return b.(int)*(index+1) + a.(int) }, 14},
		{"concat with aggregate", "", []any{"a", "b", "c"}, func(a, b any, _ int) any { return a.(string) + b.(string) }, "abc"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			s := newSeries("number", testCase.value, IntType)
			got := Agg(s, testCase.initialValue, testCase.fn)
			if got != testCase.expected {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
