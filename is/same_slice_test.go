package is

import (
	"testing"
)

func TestIs_SameSlice(t *testing.T) {
	useCases := []struct {
		name     string
		value1   []any
		value2   []any
		expected bool
	}{
		{"same slice of int", []any{1, 2, 3}, []any{1, 2, 3}, true},
		{"different slice of int", []any{1, 2, 3}, []any{1, 2, 4}, false},

		{"different slice length", []any{1, 2, 3}, []any{1, 2}, false},
		{"different slice type", []any{1, 2, 3}, []any{"1", "2", "3"}, false},
		{"nil slice", nil, nil, true},

		{"same slice of string", []any{"a", "b", "c"}, []any{"a", "b", "c"}, true},
		{"different slice of string", []any{"a", "b", "c"}, []any{"a", "b", "d"}, false},
	}

	for _, useCase := range useCases {
		t.Run(useCase.name, func(t *testing.T) {
			got := SameSlice(useCase.value1, useCase.value2)
			if got != useCase.expected {
				t.Errorf("Expected %v, got %v", useCase.expected, got)
			}
		})
	}
}
