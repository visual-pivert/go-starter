package df

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/visual-pivert/go-starter/is"
	"github.com/visual-pivert/go-starter/series"
)

// helpers
func makeDF(cols [][]any, types []string, headers []string) *Dataframe {
	ss := make([]series.Series[any], len(cols))
	for i, c := range cols {
		ss[i] = series.New(c, types[i])
	}
	return New(ss, headers)
}

func getCol(df *Dataframe, idx int) []any {
	s, _ := df.GetSeries(idx)
	return s.ToSlice()
}

func sliceEqualAny(a, b []any) bool { return reflect.DeepEqual(a, b) }

func TestDf_Shape(t *testing.T) {
	testCases := []struct {
		name     string
		cols     [][]any
		headers  []string
		types    []string
		expected []int
	}{
		{
			name:     "3x2",
			cols:     [][]any{{1, 2, 3}, {"a", "b", "c"}},
			headers:  []string{"number", "str"},
			types:    []string{"number", "string"},
			expected: []int{3, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			df := makeDF(tc.cols, tc.types, tc.headers)
			if !is.SameSlice(df.Shape(), tc.expected) {
				t.Fatalf("Shape mismatch: got %v, expected %v", df.Shape(), tc.expected)
			}
		})
	}
}

func TestDf_RemoveColumns(t *testing.T) {
	testCases := []struct {
		name            string
		cols            [][]any
		types           []string
		headers         []string
		removeIdx       []int
		expectedHeaders []string
		expectedCols    [][]any
	}{
		{
			name:            "remove single middle",
			cols:            [][]any{{1, 2, 3}, {"a", "b", "c"}, {true, false, true}},
			types:           []string{"number", "string", "bool"},
			headers:         []string{"num", "str", "flag"},
			removeIdx:       []int{1},
			expectedHeaders: []string{"num", "flag"},
			expectedCols:    [][]any{{1, 2, 3}, {true, false, true}},
		},
		{
			name:            "remove multiple out-of-order",
			cols:            [][]any{{1, 2}, {"a", "b"}, {true, false}, {10, 20}},
			types:           []string{"number", "string", "bool", "number"},
			headers:         []string{"n1", "s", "b", "n2"},
			removeIdx:       []int{3, 1},
			expectedHeaders: []string{"n1", "b"},
			expectedCols:    [][]any{{1, 2}, {true, false}},
		},
		{
			name:            "remove none (empty idx)",
			cols:            [][]any{{1}, {"a"}},
			types:           []string{"number", "string"},
			headers:         []string{"n", "s"},
			removeIdx:       []int{},
			expectedHeaders: []string{"n", "s"},
			expectedCols:    [][]any{{1}, {"a"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			df := makeDF(tc.cols, tc.types, tc.headers)
			df.RemoveColumns(tc.removeIdx)

			if !is.SameSlice(df.Shape(), []int{len(tc.cols[0]), len(tc.expectedCols)}) {
				t.Fatalf("shape after RemoveColumns: got %v, expected rows=%d cols=%d", df.Shape(), len(tc.cols[0]), len(tc.expectedCols))
			}
			if !is.SameSlice(df.headers, tc.expectedHeaders) {
				t.Fatalf("headers mismatch: got %v, expected %v", df.headers, tc.expectedHeaders)
			}
			for i := range tc.expectedCols {
				if !sliceEqualAny(getCol(df, i), tc.expectedCols[i]) {
					t.Fatalf("col %d mismatch: got %v, expected %v", i, getCol(df, i), tc.expectedCols[i])
				}
			}
		})
	}
}

func TestDf_RemoveColumnsByHeaders(t *testing.T) {
	testCases := []struct {
		name            string
		cols            [][]any
		types           []string
		headers         []string
		removeHdrs      []string
		expectedHeaders []string
	}{
		{
			name:            "remove existing header",
			cols:            [][]any{{1, 2, 3}, {"a", "b", "c"}, {true, false, true}},
			types:           []string{"number", "string", "bool"},
			headers:         []string{"num", "str", "flag"},
			removeHdrs:      []string{"str"},
			expectedHeaders: []string{"num", "flag"},
		},
		{
			name:            "remove multiple headers",
			cols:            [][]any{{1, 2, 3}, {"a", "b", "c"}, {true, false, true}},
			types:           []string{"number", "string", "bool"},
			headers:         []string{"num", "str", "flag"},
			removeHdrs:      []string{"num", "flag"},
			expectedHeaders: []string{"str"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			df := makeDF(tc.cols, tc.types, tc.headers)
			df.RemoveColumnsByHeaders(tc.removeHdrs)
			if !is.SameSlice(df.headers, tc.expectedHeaders) {
				t.Fatalf("headers mismatch after RemoveColumnsByHeaders: got %v, expected %v", df.headers, tc.expectedHeaders)
			}
		})
	}
}

func TestDf_RemoveLines(t *testing.T) {
	testCases := []struct {
		name         string
		cols         [][]any
		types        []string
		headers      []string
		removeIdx    []int
		expectedCols [][]any
	}{
		{
			name:         "remove one middle row",
			cols:         [][]any{{1, 2, 3, 4}, {"a", "b", "c", "d"}},
			types:        []string{"number", "string"},
			headers:      []string{"num", "str"},
			removeIdx:    []int{1},
			expectedCols: [][]any{{1, 3, 4}, {"a", "c", "d"}},
		},
		{
			name:         "ignore out-of-range indices",
			cols:         [][]any{{1, 2, 3}, {"a", "b", "c"}},
			types:        []string{"number", "string"},
			headers:      []string{"n", "s"},
			removeIdx:    []int{-1, 3, 99},
			expectedCols: [][]any{{1, 2, 3}, {"a", "b", "c"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			df := makeDF(tc.cols, tc.types, tc.headers)
			df.RemoveLines(tc.removeIdx)
			for i := range tc.expectedCols {
				if !sliceEqualAny(getCol(df, i), tc.expectedCols[i]) {
					t.Fatalf("unexpected col %d after RemoveLines: got %v, expected %v", i, getCol(df, i), tc.expectedCols[i])
				}
			}
		})
	}
}

func TestDf_ApplyFromBoolStatement(t *testing.T) {
	testCases := []struct {
		name         string
		cols         [][]any
		types        []string
		headers      []string
		mask         []bool
		expectedCols [][]any
	}{
		{
			name:         "filter alternating",
			cols:         [][]any{{1, 2, 3, 4}, {"a", "b", "c", "d"}},
			types:        []string{"number", "string"},
			headers:      []string{"num", "str"},
			mask:         []bool{true, false, true, false},
			expectedCols: [][]any{{1, 3}, {"a", "c"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			df := makeDF(tc.cols, tc.types, tc.headers)
			df.ApplyFromBoolStatement(series.New(tc.mask, "bool"))
			for i := range tc.expectedCols {
				if !sliceEqualAny(getCol(df, i), tc.expectedCols[i]) {
					t.Fatalf("unexpected col %d after mask: got %v, expected %v", i, getCol(df, i), tc.expectedCols[i])
				}
			}
		})
	}
}

func TestDf_ApplyFromOrderStatement(t *testing.T) {
	testCases := []struct {
		name         string
		cols         [][]any
		types        []string
		headers      []string
		order        []int
		expectedCols [][]any
	}{
		{
			name:         "permute rows",
			cols:         [][]any{{10, 20, 30}, {"x", "y", "z"}},
			types:        []string{"number", "string"},
			headers:      []string{"num", "str"},
			order:        []int{2, 0, 1},
			expectedCols: [][]any{{30, 10, 20}, {"z", "x", "y"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			df := makeDF(tc.cols, tc.types, tc.headers)
			df.ApplyFromOrderStatement(series.New(tc.order, "number"))
			for i := range tc.expectedCols {
				if !sliceEqualAny(getCol(df, i), tc.expectedCols[i]) {
					t.Fatalf("unexpected col %d after order: got %v, expected %v", i, getCol(df, i), tc.expectedCols[i])
				}
			}
		})
	}
}

func TestDf_Compute(t *testing.T) {
	testCases := []struct {
		name       string
		cols       [][]any
		types      []string
		headers    []string
		resultType string
		compute    func(d *Dataframe, idx int) any
		expected   []any
	}{
		{
			name:       "sum two ints",
			cols:       [][]any{{1, 2, 3}, {4, 5, 6}},
			types:      []string{"number", "number"},
			headers:    []string{"a", "b"},
			resultType: "number",
			compute: func(d *Dataframe, idx int) any {
				s0, _ := d.GetSeries(0)
				s1, _ := d.GetSeries(1)
				return s0.GetValue(idx).(int) + s1.GetValue(idx).(int)
			},
			expected: []any{5, 7, 9},
		},
		{
			name:       "concat string and int",
			cols:       [][]any{{"x", "y"}, {1, 2}},
			types:      []string{"string", "number"},
			headers:    []string{"s", "n"},
			resultType: "string",
			compute: func(d *Dataframe, idx int) any {
				s0, _ := d.GetSeries(0)
				s1, _ := d.GetSeries(1)
				return fmt.Sprintf("%s-%d", s0.GetValue(idx).(string), s1.GetValue(idx).(int))
			},
			expected: []any{"x-1", "y-2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			df := makeDF(tc.cols, tc.types, tc.headers)
			df.Debug()
			res := df.Compute(tc.resultType, tc.compute)
			if !reflect.DeepEqual(res.ToSlice(), tc.expected) {
				t.Fatalf("compute result mismatch: got %v, expected %v", res.ToSlice(), tc.expected)
			}
		})
	}
}
