package df

import (
	"testing"

	"github.com/visual-pivert/go-starter/is"
	"github.com/visual-pivert/go-starter/series"
)

func TestDf_Perform(t *testing.T) {
	exampleDf := createDf()
	testCases := []struct {
		name     string
		df       *Df
		fn       func(*Df, any, int) any
		expected []any
	}{
		{"int + float * 2", exampleDf, func(df *Df, prev any, idx int) any {
			return (float64(df.GetValue(1, idx).(int)) + df.GetValue(2, idx).(float64)) * 2
		}, []any{4.2, 8.4, 12.6}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			newSeries := series.NewSeries("result", []any{}, series.FloatType)
			got := Perform(tc.df, newSeries, tc.fn)
			if !is.SameSlice(tc.expected, got.GetSlice()) {
				tt.Errorf("Expected %v, got %v", tc.expected, got.GetSlice())
			}
		})
	}
}
