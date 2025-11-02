package fn

import (
	"testing"

	"github.com/visual-pivert/go-starter/is"
)

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

func TestFilterI(t *testing.T) {
	testCases := []struct {
		name     string
		value    []any
		fn       func(any) bool
		expected []int
	}{
		{"filterI even number", []any{1, 2, 3}, func(value any) bool { return value.(int)%2 == 0 }, []int{1}},
		{"filterI odd number", []any{1, 2, 3}, func(value any) bool { return value.(int)%2 == 1 }, []int{0, 2}},
		{"filterI 0", []any{1, 2, 3}, func(value any) bool { return value.(int) == 0 }, []int{}},
		{"filterI length name > 2", []any{"qwe", "qw", "q"}, func(value any) bool { return len(value.(string)) > 2 }, []int{0}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := FilterI(testCase.value, testCase.fn)
			if SameSlice(got, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}

		})

	}
}

func TestFilterTruthy(t *testing.T) {
	testCases := []struct {
		name     string
		value    []any
		expected []any
	}{
		{"filter truthy", []any{1, 0, "", "a", false}, []any{1, "a"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := FilterTruthy(testCase.value)
			if SameSlice(got, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestFilterITruthy(t *testing.T) {
	testCases := []struct {
		name     string
		value    []any
		expected []int
	}{
		{"filterI truthy", []any{1, 0, "", "a", false}, []int{0, 3}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := FilterITruthy(testCase.value)
			if SameSlice(got, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestFilterToBoolStatement(t *testing.T) {
	testCases := []struct {
		name     string
		value    []any
		expected []bool
	}{
		{"filter truthy to bool statement", []any{1, 0, "", "a", false}, []bool{true, false, false, true, false}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := FilterToBoolStatement(testCase.value, func(value any) bool {
				return is.Truthy(value)
			})
			if SameSlice(got, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
