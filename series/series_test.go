package series

import (
	"testing"

	"github.com/visual-pivert/go-starter/is"
)

func TestSeries_Append(t *testing.T) {
	testCases := []struct {
		name        string
		header      string
		stype       SeriesType
		value       []any
		expected    []any
		expectedLen int
	}{
		{"append int value to int series", "number", IntType, []any{1, 2, 3}, []any{1, 2, 3, 4}, 4},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			series := newSeries(testCase.header, testCase.value, testCase.stype)
			got := series.Append(4)
			if is.SameSlice(got.data, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
			if got.Len() != testCase.expectedLen {
				tt.Errorf("Expected %v, got %v", testCase.expectedLen, got.Len())
			}
		})
	}
}

func TestSeries_AppendSeries(t *testing.T) {
	testCases := []struct {
		name        string
		header      string
		stype       SeriesType
		value       []any
		expected    []any
		expectedLen int
	}{
		{"append int slice to int series", "number", IntType, []any{1, 2, 3}, []any{1, 2, 3, 4, 5}, 5},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			series := newSeries(testCase.header, testCase.value, testCase.stype)
			got := series.AppendSlice([]any{4, 5})
			if is.SameSlice(got.data, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
			if got.Len() != testCase.expectedLen {
				tt.Errorf("Expected %v, got %v", testCase.expectedLen, got.Len())
			}
		})
	}
}

func TestSeries_Rename(t *testing.T) {
	testCases := []struct {
		name      string
		header    string
		stype     SeriesType
		newHeader string
		expected  string
	}{
		{"rename series", "number", IntType, "new_number", "new_number"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			series := newSeries(testCase.header, []any{1, 2, 3}, testCase.stype)
			got := series.Rename(testCase.newHeader)
			if (got.Name() == testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestSeries_Set(t *testing.T) {
	testCases := []struct {
		name       string
		header     string
		stype      SeriesType
		value      []any
		indexToSet int
		newValue   any
		expected   []any
	}{
		{"set value to series", "number", IntType, []any{1, 2, 3}, 1, 4, []any{1, 4, 3}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			series := newSeries(testCase.header, testCase.value, testCase.stype)
			got := series.Set(testCase.indexToSet, testCase.newValue)
			if is.SameSlice(got.data, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestSeries_Get(t *testing.T) {
	testCases := []struct {
		name       string
		header     string
		stype      SeriesType
		value      []any
		indexToGet int
		expected   any
	}{
		{"get value from series", "number", IntType, []any{1, 2, 3}, 1, 2},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			series := newSeries(testCase.header, testCase.value, testCase.stype)
			got := series.Get(testCase.indexToGet)
			if got != testCase.expected {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestSeries_FilerToBoolStatement(t *testing.T) {
	testCases := []struct {
		name     string
		header   string
		stype    SeriesType
		fn       func(any) bool
		value    []any
		expected []bool
	}{
		{"filter even number to bool statement", "header", IntType, func(value any) bool { return value.(int)%2 == 0 }, []any{1, 2, 3}, []bool{false, true, false}},
		{"filter odd number to bool statement", "header", IntType, func(value any) bool { return value.(int)%2 == 1 }, []any{1, 2, 3}, []bool{true, false, true}},
		{"filter falsy to bool statement", "header", IntType, func(value any) bool { return is.Falsy(value) }, []any{1, 0, "", "a", false}, []bool{false, true, true, false, true}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			series := newSeries(testCase.header, testCase.value, testCase.stype)
			got := series.FilerToBoolStatement(testCase.fn)
			if is.SameSlice(got, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestSeries_IntersectWithBoolStatement(t *testing.T) {
	testCases := []struct {
		name          string
		header        string
		stype         SeriesType
		value         []any
		boolStatement []bool
		expected      []any
		expectedLen   int
	}{
		{"intersect with even bool statement", "header", IntType, []any{1, 2, 3}, []bool{false, true, false}, []any{2}, 1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			series := newSeries(testCase.header, testCase.value, testCase.stype)
			got := series.IntersectWithBoolStatement(testCase.boolStatement)
			if is.SameSlice(got.data, testCase.expected) == false {
				t.Errorf("Expected %v, got %v", testCase.expected, got)
			}
			if got.Len() != testCase.expectedLen {
				t.Errorf("Expected %v, got %v", testCase.expectedLen, got.Len())
			}
		})
	}
}

func TestSeries_ApplyWithBoolStatement(t *testing.T) {
	testCases := []struct {
		name          string
		header        string
		stype         SeriesType
		value         []any
		fn            func(any) any
		boolStatement []bool
		expected      []any
	}{
		{"apply multiply by 2 when even", "header", IntType, []any{1, 2, 3}, func(value any) any { return value.(int) * 2 }, []bool{false, true, false}, []any{1, 4, 3}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			series := newSeries(testCase.header, testCase.value, testCase.stype)
			got := series.ApplyWithBoolStatement(testCase.boolStatement, testCase.fn)
			if is.SameSlice(got.data, testCase.expected) == false {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
