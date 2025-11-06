package df

import (
	"github.com/visual-pivert/go-starter/series"
)

func Compute(dataframe *Df, s *series.Series, fun func(d *Df, prev any, index int) any) *series.Series {
	newSeries := series.New(s.Name(), []any{}, s.Type())
	shape := dataframe.Shape()
	nbLine := shape[0]
	var prev any
	for i := 0; i < nbLine; i++ {
		value := fun(dataframe, prev, i)
		newSeries = newSeries.Append(value)
		prev = value
	}
	return newSeries
}
