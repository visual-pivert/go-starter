package fn

import (
	"testing"

	"github.com/visual-pivert/go-starter/is"
)

func TestReduce(t *testing.T) {
	testCases := []struct {
		name         string
		initialValue any
		value        []any
		fn           func(any, any) any
		expected     []any
	}{
		{"cumsum", 0, []any{1, 2, 3}, func(a, b any) any { return a.(int) + b.(int) }, []any{1, 3, 6}},
		{"cummultiply", 1, []any{1, 2, 3}, func(a, b any) any { return a.(int) * b.(int) }, []any{1, 2, 6}},
		{"cumconcat", "", []any{"a", "b", "c"}, func(a, b any) any { return a.(string) + b.(string) }, []any{"a", "ab", "abc"}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := Reduce(testCase.value, testCase.initialValue, testCase.fn)
			if is.SameSlice(got, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
