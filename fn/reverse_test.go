package fn

import (
	"testing"

	"github.com/visual-pivert/go-starter/is"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		value    []any
		expected []any
	}{
		{"reverse", []any{"a", "b", "c"}, []any{"c", "b", "a"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := Reverse(testCase.value)
			if is.SameSlice(got, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
