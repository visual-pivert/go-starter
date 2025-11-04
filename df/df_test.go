package df

import (
	"testing"

	"github.com/visual-pivert/go-starter/is"
	"github.com/visual-pivert/go-starter/series"
)

func createDf() *Df {
	dataframe := NewDf(
		series.NewSeries("str", []any{"a", "b", "c"}, series.StringType),
		series.NewSeries("int", []any{1, 2, 3}, series.IntType),
		series.NewSeries("float", []any{1.1, 2.2, 3.3}, series.FloatType),
		series.NewSeries("bool", []any{true, false, true}, series.BoolType),
		series.NewSeries("time", []any{"2025/10/04", "2025/11/11", "2025/12/07"}, series.TimeType),
	)
	return dataframe
}

func TestDf_Types(t *testing.T) {
	testCases := []struct {
		name           string
		df             *Df
		expected       []series.SeriesType
		expectedString []string
	}{
		{"types", createDf(), []series.SeriesType{series.StringType, series.IntType, series.FloatType, series.BoolType, series.TimeType}, []string{"string", "int", "float", "bool", "time"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := testCase.df.Types()
			gotStr := testCase.df.TypesToString()
			if !is.SameSlice(got, testCase.expected) {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
			if !is.SameSlice(gotStr, testCase.expectedString) {
				tt.Errorf("Expected %v, got %v", testCase.expectedString, gotStr)
			}
		})
	}
}

func TestDf_GetSeries(t *testing.T) {
	testCases := []struct {
		name               string
		df                 *Df
		idx                int
		expectedSeriesName string
	}{
		{"get str series", createDf(), 0, "str"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.df.GetSeries(testCase.idx)
			if got.Name() != testCase.expectedSeriesName {
				t.Errorf("Expected %v, got %v", testCase.expectedSeriesName, got.Name())
			}
		})
	}
}

func TestDf_GetSeriesByHeader(t *testing.T) {
	testCases := []struct {
		name               string
		df                 *Df
		headerName         string
		expectedSeriesName string
	}{
		{"get str series by header", createDf(), "str", "str"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.df.GetSeriesByHeader(testCase.headerName)
			if got.Name() != testCase.expectedSeriesName {
				t.Errorf("Expected %v, got %v", testCase.expectedSeriesName, got.Name())
			}
		})
	}
}

func TestDf_RemoveSeries(t *testing.T) {
	testCases := []struct {
		name            string
		df              *Df
		idx             []int
		expectedLen     int
		expectedColumns []string
	}{
		{"Remove str series", createDf(), []int{0}, 4, []string{"int", "float", "bool", "time"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := testCase.df.RemoveColumns(testCase.idx...)
			testCase.df.Debug()
			got.Debug()
			if got.Shape()[1] != testCase.expectedLen {
				tt.Errorf("Expected %v, got %v", testCase.expectedLen, got.Shape()[1])
			}
			if !is.SameSlice(got.columns, testCase.expectedColumns) {
				tt.Errorf("Expected %v got %v", testCase.expectedColumns, got.columns)
			}
		})
	}
}

func TestDf_RemoveSeriesByHeader(t *testing.T) {
	testCases := []struct {
		name            string
		df              *Df
		headers         []string
		expectedLen     int
		expectedColumns []string
	}{
		{"Remove str series", createDf(), []string{"str"}, 4, []string{"int", "float", "bool", "time"}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := testCase.df.RemoveColumnsByHeader(testCase.headers...)
			if got.Shape()[1] != testCase.expectedLen {
				tt.Errorf("Expected %v, got %v", testCase.expectedLen, got.Shape()[1])
			}
			if !is.SameSlice(got.columns, testCase.expectedColumns) {
				tt.Errorf("Expected %v got %v", testCase.expectedColumns, got.columns)
			}
		})
	}
}

func TestDf_Shape(t *testing.T) {
	testCases := []struct {
		name     string
		df       *Df
		expected []int
	}{
		{"shape", createDf(), []int{3, 5}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := testCase.df.Shape()
			if !is.SameSlice(got, testCase.expected) {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}

		})
	}
}

func TestDf_IntersectWithBoolStatement(t *testing.T) {
	testCases := []struct {
		name          string
		df            *Df
		boolStatement []bool
		expectedShape []int
	}{
		{"intersect with bool statement", createDf(), []bool{true, false, true}, []int{2, 5}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := testCase.df.IntersectWithBoolStatement(testCase.boolStatement)
			if !is.SameSlice(got.Shape(), testCase.expectedShape) {
				tt.Errorf("Expected %v, got %v", testCase.expectedShape, got.Shape())
			}
		})
	}
}
