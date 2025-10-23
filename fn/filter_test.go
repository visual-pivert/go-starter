package fn

import "testing"

func TestFilter(t *testing.T) {
	testCases := []struct {
		name     string
		value    []any
		fn       func(any) bool
		expected []any
	}{
		{"filter even number", []any{1, 2, 3}, func(value any) bool { return value.(int)%2 == 0 }, []any{2}},
		{"filter odd number", []any{1, 2, 3}, func(value any) bool { return value.(int)%2 == 1 }, []any{1, 3}},
		{"filter 0", []any{1, 2, 3}, func(value any) bool { return value.(int) == 0 }, []any{}},
		{"filter length name > 2", []any{"qwe", "qw", "q"}, func(value any) bool { return len(value.(string)) > 2 }, []any{"qwe"}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := Filter(testCase.value, testCase.fn)
			if SameSlice(got, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}

		})

	}
}
