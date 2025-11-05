package extract

import (
	"testing"

	"github.com/visual-pivert/go-starter/fn"
	"github.com/visual-pivert/go-starter/is"
	"github.com/visual-pivert/go-starter/series"
)

func TestToDf(t *testing.T) {
	testCases := []struct {
		name            string
		value           [][]string
		expectedHeaders []string
		expectedTypes   []series.Type
		expectedLen     int
	}{
		{"to df", [][]string{{"letters", "numbers"}, {"a", "1"}, {"b", "2"}, {"c", "3"}}, []string{"letters", "numbers"}, []series.Type{series.StringType, series.IntType}, 2},
		{"to df 2", [][]string{{"times", "numbers"}, {"2025/10/01", "1"}, {"2024/10/10", "2"}, {"2025/11/12", "3"}}, []string{"times", "numbers"}, []series.Type{series.TimeType, series.IntType}, 2},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := ToDf(testCase.value, 0)
			gotHeaders := got.Columns()
			gotTypes := fn.Map(got.GetAllSeries(), func(t *series.Series, idx int) series.Type {
				return t.Type()
			})
			if is.SameSlice(gotHeaders, testCase.expectedHeaders) == false {
				tt.Errorf("Expected %v, got %v", testCase.expectedHeaders, gotHeaders)
			}
			if is.SameSlice(gotTypes, testCase.expectedTypes) == false {
				tt.Errorf("Expected %v, got %v", testCase.expectedTypes, gotTypes)
			}
		})
	}
}
