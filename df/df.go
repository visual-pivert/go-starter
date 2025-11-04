package df

import (
	"math"

	"github.com/visual-pivert/go-starter/fn"
	"github.com/visual-pivert/go-starter/series"
)

type Df struct {
	data    []*series.Series
	columns []string
	shape   []int
}

func calculateColumns(data []*series.Series) []string {
	var out []string
	for _, s := range data {
		out = append(out, s.Name())
	}
	return out
}

func seriesMaxLen(data []*series.Series) int {
	maxValue := 0
	for _, s := range data {
		maxValue = int(math.Max(float64(s.Len()), float64(maxValue)))
	}
	return maxValue
}

func calculateShape(data []*series.Series) []int {
	out := []int{0, 0}
	out[1] = len(data)
	out[0] = seriesMaxLen(data)
	return out
}

func NewDf(data ...*series.Series) *Df {
	return &Df{data, calculateColumns(data), calculateShape(data)}
}

func (dataframe *Df) Columns() []string {
	return dataframe.columns
}

func (dataframe *Df) GetSeries(idx int) *series.Series {
	return dataframe.data[idx]
}

func (dataframe *Df) GetSeriesByHeader(header string) *series.Series {
	idx := fn.IndexOf(header, dataframe.columns)
	return dataframe.data[idx]
}

func (dataframe *Df) GetValue(column int, row int) any {
	return dataframe.data[column].Get(row)
}

func (dataframe *Df) GetValueByHeader(column string, row int) any {
	idx := fn.IndexOf(column, dataframe.columns)
	return dataframe.data[idx].Get(row)
}

func (dataframe *Df) Shape() []int {
	return dataframe.shape
}

func (dataframe *Df) RemoveColumns(indexes ...int) *Df {
	var newSeries []*series.Series
	for i, d := range dataframe.data {
		index := fn.IndexOf(i, indexes)
		if index == -1 {
			newSeries = append(newSeries, d)
		}
	}
	return NewDf(newSeries...)
}

func (dataframe *Df) RemoveColumnsByHeader(names ...string) *Df {
	var newSeries []*series.Series
	for _, d := range dataframe.data {
		index := fn.IndexOf(d.Name(), names)
		if index == -1 {
			newSeries = append(newSeries, d)
		}
	}
	return NewDf(newSeries...)
}
