package series

import (
	"testing"

	"github.com/visual-pivert/go-starter/is"
)

func TestSeries_Append(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		append   []int
		expected []int
		lenExp   int
	}{
		{"append single", []int{1, 2, 3}, "number", []int{4}, []int{1, 2, 3, 4}, 4},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.Append(tc.append)
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
			if got.Len() != tc.lenExp {
				tt.Errorf("Expected %v, got %v", tc.lenExp, got.Len())
			}
		})
	}
}

func TestSeries_AppendTo(t *testing.T) {
	testCases := []struct {
		name     string
		value    []string
		t        string
		pos      int
		append   []string
		expected []string
	}{
		{"insert in middle", []string{"a", "c"}, "number", 1, []string{"b"}, []string{"a", "b", "c"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.AppendTo(tc.pos, tc.append)
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
		})
	}
}

func TestSeries_Pop(t *testing.T) {
	// pop
	testCases := []struct {
		name     string
		value    []int
		t        string
		expected []int
		last     int
	}{
		{"pop last", []int{1, 2, 3}, "number", []int{1, 2}, 3},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, "string")
			after, last := s.Pop()
			if last != tc.last {
				tt.Errorf("Expected %v, got %v", tc.last, last)
			}
			if !is.SameSlice(after.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, after.ToSlice())
			}
		})
	}
}

func TestSeries_Shift(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		expected []int
		first    int
	}{
		{"shift first", []int{1, 2, 3}, "number", []int{2, 3}, 1},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			after, first := s.Shift()
			if first != tc.first {
				tt.Errorf("Expected %v, got %v", tc.first, first)
			}
			if !is.SameSlice(after.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, after.ToSlice())
			}
		})
	}
}

func TestSeries_Remove(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		index    int
		expected []int
	}{
		{"remove index 1", []int{1, 2, 3, 4}, "number", 1, []int{1, 3, 4}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.Remove(tc.index)
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
		})
	}
}

func TestSeries_Range(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		start    int
		nbr      int
		expected []int
	}{
		{"slice middle", []int{1, 2, 3, 4}, "number", 1, 2, []int{2, 3}},
		{"slice last", []int{1, 2, 3, 4}, "number", 1, 3, []int{2, 3, 4}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.Range(tc.start, tc.nbr)
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
		})
	}
}

func TestSeries_Len_Count(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		lenExp   int
		countExp int
		sliceExp []int
	}{
		{"length", []int{10, 20, 30}, "number", 3, 3, []int{10, 20, 30}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			if s.Len() != tc.lenExp || s.Count() != tc.countExp {
				tt.Errorf("Expected length %d, got %d/%d", tc.lenExp, s.Len(), s.Count())
			}
		})
	}
}

// TODO: expected slice is not []int but []any and we need to test type
func TestSeries_ToSlice(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		lenExp   int
		countExp int
		sliceExp []int
	}{
		{"slice", []int{10, 20, 30}, "number", 3, 3, []int{10, 20, 30}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			if !is.SameSlice(s.ToSlice(), tc.sliceExp) {
				tt.Errorf("Expected %v, got %v", tc.sliceExp, s.ToSlice())
			}
		})
	}
}

func TestSeries_Get(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		index    int
		expected int
	}{
		{"get index 1", []int{1, 2, 3}, "number", 1, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.GetValue(tc.index)
			if got != tc.expected {
				tt.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestSeries_Set(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		index    int
		newValue int
		expected []int
	}{
		{"set index 1", []int{1, 2, 3}, "number", 1, 42, []int{1, 42, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.SetValue(tc.index, tc.newValue)
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
		})
	}
}

func TestSeries_Filter(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		fn       func(int) bool
		expected []int
	}{
		{"even values", []int{1, 2, 3, 4, 5}, "number", func(v int) bool { return v%2 == 0 }, []int{2, 4}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.Filter(tc.fn)
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
		})
	}
}

func TestSeries_FilterI(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		fn       func(int) bool
		expected []int
	}{
		{"even indices", []int{1, 2, 3, 4, 5}, "number", func(v int) bool { return v%2 == 0 }, []int{1, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			idx := s.FilterI(tc.fn)
			if !is.SameSlice(idx.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, idx.ToSlice())
			}
		})
	}
}

func TestSeries_Map(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		fn       func(int, int) int
		expected []int
	}{
		{"double", []int{1, 2, 3}, "number", func(v int, i int) int { return v * 2 }, []int{2, 4, 6}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.Map(tc.fn)
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
		})
	}
}

func TestSeries_MapToBool(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		fn       func(int, int) bool
		expected []bool
	}{
		{"odd to bool", []int{1, 2, 3}, "number", func(v int, i int) bool { return v%2 == 1 }, []bool{true, false, true}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.MapToBool(tc.fn)
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
		})
	}
}

func TestSeries_Reduce(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		initial  int
		fn       func(int, int, int) int
		expected []int
	}{
		{"cumulative sum", []int{1, 2, 3, 4}, "number", 0, func(last int, curr int, idx int) int { return last + curr }, []int{1, 3, 6, 10}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.Reduce(tc.initial, tc.fn)
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
		})
	}

}

func TestSeries_Agg(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		t        string
		initial  int
		fn       func(int, int, int) int
		expected int
	}{
		{"aggregate sum", []int{1, 2, 3, 4}, "number", 0, func(last int, curr int, idx int) int { return last + curr }, 10},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.Agg(tc.initial, tc.fn)
			if got != tc.expected {
				tt.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestSeries_Any(t *testing.T) {
	testCases := []struct {
		name     string
		value    []string
		t        string
		fn       func(string) bool
		expected bool
	}{
		{"any length > 2", []string{"a", "bb", "ccc"}, "string", func(v string) bool { return len(v) > 2 }, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.Any(tc.fn)
			if got != tc.expected {
				tt.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}

}

func TestSeries_All(t *testing.T) {
	testCases := []struct {
		name     string
		value    []string
		t        string
		fn       func(string) bool
		expected bool
	}{
		{"all length >=1", []string{"a", "bb", "ccc"}, "string", func(v string) bool { return len(v) >= 1 }, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.All(tc.fn)
			if got != tc.expected {
				tt.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestSeries_IndexOf(t *testing.T) {
	testCases := []struct {
		name     string
		value    []string
		t        string
		lookFor  string
		expected int
	}{
		{"index of bb", []string{"a", "bb", "ccc"}, "string", "bb", 1},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.IndexOf(tc.lookFor)
			if got != tc.expected {
				tt.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestSeries_Reverse(t *testing.T) {
	testCases := []struct {
		name     string
		value    []string
		t        string
		expected []string
	}{
		{"reverse order", []string{"a", "bb", "ccc"}, "string", []string{"ccc", "bb", "a"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value, tc.t)
			got := s.Reverse()
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
		})
	}
}

func TestSeries_ApplyBoolStatement(t *testing.T) {
	testCases := []struct {
		name          string
		value         []string
		t             string
		boolStatement []bool
		expected      []string
	}{
		{"apply bool statement", []string{"a", "bb", "ccc"}, "string", []bool{true, false, true}, []string{"a", "ccc"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			seriesValue := New(testCase.value, testCase.t)
			seriesBool := New(testCase.boolStatement, "bool")
			got := seriesValue.ApplyBoolStatement(seriesBool)
			if !is.SameSlice(got.ToSlice(), testCase.expected) {
				tt.Errorf("Expected %v, got %v", testCase.expected, got.ToSlice())
			}
		})
	}
}

func TestSeries_ApplyOrderStatement(t *testing.T) {
	testCases := []struct {
		name           string
		value          []string
		t              string
		orderStatement []int
		expected       []string
	}{
		{"apply order statement", []string{"a", "bb", "ccc"}, "string", []int{2, 0, 1}, []string{"ccc", "a", "bb"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			seriesValue := New(testCase.value, testCase.t)
			seriesOrder := New(testCase.orderStatement, testCase.t)
			got := seriesValue.ApplyOrderStatement(seriesOrder)
			if !is.SameSlice(got.ToSlice(), testCase.expected) {
				tt.Errorf("Expected %v, got %v", testCase.expected, got.ToSlice())
			}
		})
	}
}

func TestSeries_CountValue(t *testing.T) {
	testCases := []struct {
		name         string
		value        []any
		t            string
		valueToCount any
		expected     int
	}{
		{"count values", []any{1, 2, 2, 3}, "number", 2, 2},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			seriesValue := New(testCase.value, testCase.t)
			seriesValue.Debug()
			got := seriesValue.CountValue(2)
			if got != testCase.expected {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
