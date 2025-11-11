package series

import (
	"testing"

	"github.com/visual-pivert/go-starter/is"
)

func TestSeries_Append(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		append   []int
		expected []int
		lenExp   int
	}{
		{"append single", []int{1, 2, 3}, []int{4}, []int{1, 2, 3, 4}, 4},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		pos      int
		append   []string
		expected []string
	}{
		{"insert in middle", []string{"a", "c"}, 1, []string{"b"}, []string{"a", "b", "c"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		expected []int
		last     int
	}{
		{"pop last", []int{1, 2, 3}, []int{1, 2}, 3},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		expected []int
		first    int
	}{
		{"shift first", []int{1, 2, 3}, []int{2, 3}, 1},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		index    int
		expected []int
	}{
		{"remove index 1", []int{1, 2, 3, 4}, 1, []int{1, 3, 4}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		start    int
		nbr      int
		expected []int
	}{
		{"slice middle", []int{1, 2, 3, 4}, 1, 2, []int{2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		lenExp   int
		countExp int
		sliceExp []int
	}{
		{"length", []int{10, 20, 30}, 3, 3, []int{10, 20, 30}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
			if s.Len() != tc.lenExp || s.Count() != tc.countExp {
				tt.Errorf("Expected length %d, got %d/%d", tc.lenExp, s.Len(), s.Count())
			}
		})
	}
}
func TestSeries_ToSlice(t *testing.T) {
	testCases := []struct {
		name     string
		value    []int
		lenExp   int
		countExp int
		sliceExp []int
	}{
		{"slice", []int{10, 20, 30}, 3, 3, []int{10, 20, 30}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		index    int
		expected int
	}{
		{"get index 1", []int{1, 2, 3}, 1, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		index    int
		newValue int
		expected []int
	}{
		{"set index 1", []int{1, 2, 3}, 1, 42, []int{1, 42, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		fn       func(int) bool
		expected []int
	}{
		{"even values", []int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 }, []int{2, 4}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		fn       func(int) bool
		expected []int
	}{
		{"even indices", []int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 }, []int{1, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		fn       func(int, int) int
		expected []int
	}{
		{"double", []int{1, 2, 3}, func(v int, i int) int { return v * 2 }, []int{2, 4, 6}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		fn       func(int, int) bool
		expected []bool
	}{
		{"odd to bool", []int{1, 2, 3}, func(v int, i int) bool { return v%2 == 1 }, []bool{true, false, true}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		initial  int
		fn       func(int, int, int) int
		expected []int
	}{
		{"cumulative sum", []int{1, 2, 3, 4}, 0, func(last int, curr int, idx int) int { return last + curr }, []int{1, 3, 6, 10}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		initial  int
		fn       func(int, int, int) int
		expected int
	}{
		{"aggregate sum", []int{1, 2, 3, 4}, 0, func(last int, curr int, idx int) int { return last + curr }, 10},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		fn       func(string) bool
		expected bool
	}{
		{"any length > 2", []string{"a", "bb", "ccc"}, func(v string) bool { return len(v) > 2 }, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		fn       func(string) bool
		expected bool
	}{
		{"all length >=1", []string{"a", "bb", "ccc"}, func(v string) bool { return len(v) >= 1 }, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		lookFor  string
		expected int
	}{
		{"index of bb", []string{"a", "bb", "ccc"}, "bb", 1},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
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
		expected []string
	}{
		{"reverse order", []string{"a", "bb", "ccc"}, []string{"ccc", "bb", "a"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			s := New(tc.value)
			got := s.Reverse()
			if !is.SameSlice(got.ToSlice(), tc.expected) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.ToSlice())
			}
		})
	}
}
