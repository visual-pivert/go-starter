package df

import (
	"fmt"
	"math"
	"strings"

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

func (dataframe *Df) Types() []series.SeriesType {
	var out []series.SeriesType
	for _, s := range dataframe.data {
		out = append(out, s.Type())
	}
	return out
}

func (dataframe *Df) TypesToString() []string {
	t := dataframe.Types()
	var out []string
	for _, s := range t {
		out = append(out, string(s))
	}
	return out
}

func (dataframe *Df) GetSeries(idx int) *series.Series {
	return dataframe.data[idx]
}

func (dataframe *Df) GetAllSeries() []*series.Series {
	return dataframe.data
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

func (dataframe *Df) AddSeries(s *series.Series) *Df {
	return NewDf(append(dataframe.data, s)...)
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

func (dataframe *Df) ToMdString() string {
	cols := len(dataframe.columns)
	if cols == 0 {
		return ""
	}

	rows := dataframe.shape[0]

	// Compute column widths from headers and all visible values
	widths := make([]int, cols)
	for j := 0; j < cols; j++ {
		w := len(dataframe.columns[j] + "(" + dataframe.TypesToString()[j] + ")")
		col := dataframe.data[j]
		for i := 0; i < rows; i++ {
			var v any
			if i < col.Len() {
				v = col.Get(i)
			} else {
				v = ""
			}
			s := fmt.Sprint(v)
			if l := len(s); l > w {
				w = l
			}
		}
		// Markdown requires at least 3 dashes in the separator
		if w < 3 {
			w = 3
		}
		widths[j] = w
	}

	pad := func(s string, w int) string {
		if len(s) >= w {
			return s
		}
		return s + strings.Repeat(" ", w-len(s))
	}

	lines := make([]string, 0, rows+2)
	// Header row
	{
		row := "|"
		for j := 0; j < cols; j++ {
			row += " " + pad(dataframe.columns[j]+"("+dataframe.TypesToString()[j]+")", widths[j]) + " |"
		}
		lines = append(lines, row)
	}
	// Separator row (at least 3 dashes per column)
	{
		row := "|"
		for j := 0; j < cols; j++ {
			nd := widths[j]
			if nd < 3 {
				nd = 3
			}
			row += " " + strings.Repeat("-", nd) + " |"
		}
		lines = append(lines, row)
	}
	// Data rows
	for i := 0; i < rows; i++ {
		row := "|"
		for j := 0; j < cols; j++ {
			col := dataframe.data[j]
			var v any
			if i < col.Len() {
				v = col.Get(i)
			} else {
				v = ""
			}
			row += " " + pad(fmt.Sprint(v), widths[j]) + " |"
		}
		lines = append(lines, row)
	}

	return strings.Join(lines, "\n")
}

func (dataframe *Df) Debug() {
	fmt.Println(dataframe.ToMdString())
}
