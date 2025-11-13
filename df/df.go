package df

import (
	"fmt"
	"strings"

	fnVisual "github.com/visual-pivert/go-starter/fn"
	"github.com/visual-pivert/go-starter/is"
	"github.com/visual-pivert/go-starter/series"
)

type Dataframe struct {
	sheet   []series.Series[any]
	headers []string
}

func New(sheet []series.Series[any], headers []string) *Dataframe {
	return &Dataframe{
		sheet:   sheet,
		headers: headers,
	}
}

func (df *Dataframe) GetSeries(idx int) (series.Series[any], string) {
	return df.sheet[idx], df.headers[idx]
}

func (df *Dataframe) GetSeriesByHeader(name string) (series.Series[any], string) {
	idx := fnVisual.IndexOf(name, df.headers)
	return df.GetSeries(idx)
}

func (df *Dataframe) Shape() []int {
	return []int{df.sheet[0].Len(), len(df.sheet)}
}

func (df *Dataframe) Append(s series.Series[any], header string) {
	df.sheet = append(df.sheet, s)
	df.headers = append(df.headers, header)
}

func (df *Dataframe) Copy() *Dataframe {
	return New(df.sheet, df.headers)
}

func (df *Dataframe) ApplyFromBoolStatement(boolStatement series.Series[bool]) {
	newSheet := make([]series.Series[any], 0, len(df.sheet))
	for _, col := range df.sheet {
		newSheet = append(newSheet, col.ApplyBoolStatement(boolStatement))
	}
	df.sheet = newSheet
}

func (df *Dataframe) ApplyFromOrderStatement(orderStatement series.Series[int]) {
	newSheet := make([]series.Series[any], 0, len(df.sheet))
	for _, col := range df.sheet {
		newSheet = append(newSheet, col.ApplyOrderStatement(orderStatement))
	}
	df.sheet = newSheet
}

func (df *Dataframe) RemoveColumns(idx []int) {
	newSheet := make([]series.Series[any], 0, len(df.sheet)-len(idx))
	newHeaders := make([]string, 0, len(df.headers)-len(idx))

	for i := range df.sheet {
		if !is.In(i, idx) {
			newSheet = append(newSheet, df.sheet[i])
			newHeaders = append(newHeaders, df.headers[i])
		}
	}

	df.sheet = newSheet
	df.headers = newHeaders
}

func (df *Dataframe) RemoveColumnsByHeaders(headers []string) {
	idx := make([]int, len(headers))
	for i, header := range headers {
		idx[i] = fnVisual.IndexOf(header, df.headers)
	}
	df.RemoveColumns(idx)
}

func (df *Dataframe) RemoveLines(idx []int) {
	line := df.Shape()[0]
	boolStatement := make([]bool, line)
	for i := 0; i < line; i++ {
		boolStatement[i] = true
	}
	for _, i := range idx {
		if i >= 0 && i < line {
			boolStatement[i] = false
		}
	}
	boolStatementSeries := series.New(boolStatement, "bool")
	df.ApplyFromBoolStatement(boolStatementSeries)
}

func (df *Dataframe) Compute(t string, fn func(d *Dataframe, idx int) any) series.Series[any] {
	rows := df.Shape()[0]
	slice := make([]any, rows)
	s := series.New(slice, t)
	for i := 0; i < rows; i++ {
		s.SetValue(i, fn(df, i))
	}
	return s
}

func (df *Dataframe) Debug() {
	if len(df.sheet) == 0 {
		fmt.Println("(empty dataframe)")
		return
	}
	cols := len(df.sheet)
	rows := df.Shape()[0]

	// Build header with types
	hdr := make([]string, cols)
	for i, h := range df.headers {
		t := df.sheet[i].Type()
		hdr[i] = fmt.Sprintf("%s(%s)", h, t)
	}

	// Build cell strings for all rows
	cells := make([][]string, rows)
	for r := 0; r < rows; r++ {
		row := make([]string, cols)
		for c := 0; c < cols; c++ {
			row[c] = fmt.Sprintf("%v", df.sheet[c].GetValue(r))
		}
		cells[r] = row
	}

	// Compute max width per column (consider header and all rows)
	widths := make([]int, cols)
	for c := 0; c < cols; c++ {
		w := len(hdr[c])
		for r := 0; r < rows; r++ {
			if lw := len(cells[r][c]); lw > w {
				w = lw
			}
		}
		widths[c] = w
	}

	pad := func(s string, w int) string {
		if len(s) < w {
			return s + strings.Repeat(" ", w-len(s))
		}
		return s
	}

	// Print header with padding
	head := make([]string, cols)
	for c := 0; c < cols; c++ {
		head[c] = pad(hdr[c], widths[c])
	}
	fmt.Printf("| %s |\n", strings.Join(head, " | "))

	// Separator matching widths
	sep := make([]string, cols)
	for c := 0; c < cols; c++ {
		sep[c] = strings.Repeat("-", widths[c])
	}
	fmt.Printf("| %s |\n", strings.Join(sep, " | "))

	// Print rows with padding
	for r := 0; r < rows; r++ {
		row := make([]string, cols)
		for c := 0; c < cols; c++ {
			row[c] = pad(cells[r][c], widths[c])
		}
		fmt.Printf("| %s |\n", strings.Join(row, " | "))
	}
}
