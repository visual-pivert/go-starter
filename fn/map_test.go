package fn

import "testing"

func TestMap(t *testing.T) {
	useCases := []struct {
		name     string
		value    []any
		fn       func(any) any
		expected []any
	}{
		{"multiply by 2", []any{1, 2, 3}, func(value any) any { return value.(int) * 2 }, []any{2, 4, 6}},
		{"even number", []any{1, 2, 3}, func(value any) any { return value.(int)%2 == 0 }, []any{false, true, false}},
	}

	for _, useCase := range useCases {
		t.Run(useCase.name, func(t *testing.T) {
			got := Map(useCase.value, useCase.fn)
			if !SameSlice(got, useCase.expected) {
				t.Errorf("Expected %v, got %v", useCase.expected, got)
			}
		})
	}
}

func TestMapReverse(t *testing.T) {
	useCases := []struct {
		name     string
		value    []any
		fn       func(any) any
		expected []any
	}{
		{"multiply by 2 (reversed)", []any{1, 2, 3}, func(value any) any { return value.(int) * 2 }, []any{6, 4, 2}},
		{"even number (reversed)", []any{1, 2, 3}, func(value any) any { return value.(int)%2 == 0 }, []any{false, true, false}},
	}

	for _, useCase := range useCases {
		t.Run(useCase.name, func(t *testing.T) {
			got := MapReverse(useCase.value, useCase.fn)
			if !SameSlice(got, useCase.expected) {
				t.Errorf("Expected %v, got %v", useCase.expected, got)
			}
		})
	}
}
