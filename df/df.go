// Package df provides a simple in-memory dataframe made of generic series.Series columns
// and utility functions to slice, filter, reorder, and inspect tabular data.

package df

import (
	"fmt"
	"strings"

	fnVisual "github.com/visual-pivert/go-starter/fn"
	"github.com/visual-pivert/go-starter/is"
	"github.com/visual-pivert/go-starter/series"
)

// Dataframe represents a simple 2D tabular structure made of multiple Series.
// Each column is a generic Series[any] and headers keeps the column names aligned
// with the underlying sheet order.
type Dataframe struct {
	sheet   []series.Series[any]
	headers []string
}

// New creates a new Dataframe from a slice of Series and a matching list of headers.
// Preconditions:
// - All series in sheet should have identical length (same number of rows).
// - headers must have the same length as sheet.
// Notes:
// - No validation is performed; violating the preconditions may cause panics in later operations.
// Examples:
//
//	s1 := series.New([]int{1, 2, 3}, "number")
//	s2 := series.New([]string{"a", "b", "c"}, "string")
//	df := df.New([]series.Series[any]{s1, s2}, []string{"col1", "col2"})
func New(sheet []series.Series[any], headers []string) *Dataframe {
	return &Dataframe{
		sheet:   sheet,
		headers: headers,
	}
}

// GetSeries returns the Series and its header at the given zero-based column index.
// It panics if idx is out of range.
// Examples:
//
//	s1, header := df.GetSeries(0) // s1 is the Series at index 0, and header is the header at index 0
func (df *Dataframe) GetSeries(idx int) (series.Series[any], string) {
	return df.sheet[idx], df.headers[idx]
}

// GetHeaders returns the headers (column names) of the Dataframe in column order.
// The returned slice is the underlying slice; do not modify it unless you know
// what you're doing.
func (df *Dataframe) GetHeaders() []string {
	return df.headers
}

// GetSeriesByHeader returns the Series and its header by column name.
// It panics if the given name is not found (because it resolves to index -1).
// Examples:
//
//	s1, header := df.GetSeriesByHeader("col1") // s1 is the Series with header "col1", and header is "col1"
func (df *Dataframe) GetSeriesByHeader(name string) (series.Series[any], string) {
	idx := fnVisual.IndexOf(name, df.headers)
	return df.GetSeries(idx)
}

// Shape returns the shape of the Dataframe as a slice of integers: [rows, columns].
// Notes:
// - Panics if the dataframe has zero columns.
// Examples:
//
//	rows, cols := df.Shape() // rows is the number of rows, and cols is the number of columns
func (df *Dataframe) Shape() []int {
	return []int{df.sheet[0].Len(), len(df.sheet)}
}

// Append adds a Series to the Dataframe with the given header.
// Preconditions:
// - The appended series must have the same number of rows as existing columns.
// Notes:
// - No validation is performed.
// Examples:
//
//	s1 := series.New([]int{1, 2, 3}, "number")
//	df.Append(s1, "col1")
func (df *Dataframe) Append(s series.Series[any], header string) {
	df.sheet = append(df.sheet, s)
	df.headers = append(df.headers, header)
}

// Copy returns a shallow copy of the Dataframe.
// Notes:
//   - The returned Dataframe shares the underlying Series with the original
//     (no deep copy of column data is performed).
func (df *Dataframe) Copy() *Dataframe {
	return New(df.sheet, df.headers)
}

// ApplyFromBoolStatement applies a boolean mask to every column of the Dataframe.
// True keeps the row; false removes it. The operation is applied to each Series
// and the resulting columns replace the originals.
// Preconditions:
// - boolStatement.Len() must equal the number of rows in the dataframe.
// Notes:
// - No validation is performed; a length mismatch will likely panic in the underlying Series.
// Examples:
//
//	df.ApplyFromBoolStatement(series.New([]bool{true, false, true}, "bool"))
func (df *Dataframe) ApplyFromBoolStatement(boolStatement series.Series[bool]) {
	newSheet := make([]series.Series[any], 0, len(df.sheet))
	for _, col := range df.sheet {
		newSheet = append(newSheet, col.ApplyBoolStatement(boolStatement))
	}
	df.sheet = newSheet
}

// ApplyFromOrderStatement reorders all rows according to an order index series.
// Each column is reordered using the same order indices, keeping rows aligned
// across columns. The resulting columns replace the originals.
// Preconditions:
// - orderStatement.Len() must equal the number of rows.
// - Indices should be within [0, rows) and ideally form a permutation.
// Notes:
//   - No validation is performed; invalid indices or duplicates may cause panics
//     or undefined reordering in the underlying Series.
//
// Examples:
//
//	// move row 0 to the end
//	df.ApplyFromOrderStatement(series.New([]int{1, 2, 0}, "number"))
func (df *Dataframe) ApplyFromOrderStatement(orderStatement series.Series[int]) {
	newSheet := make([]series.Series[any], 0, len(df.sheet))
	for _, col := range df.sheet {
		newSheet = append(newSheet, col.ApplyOrderStatement(orderStatement))
	}
	df.sheet = newSheet
}

// RemoveColumns removes columns at the provided zero-based indices.
// Indices not present are ignored; duplicates are effectively removed once.
// Order of remaining columns is preserved.
// Examples:
//
//	df.RemoveColumns([]int{0, 2})
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

// RemoveColumnsByHeaders removes columns by their header names.
// Headers not present are ignored. Order of remaining columns is preserved.
// Examples:
//
//	df.RemoveColumnsByHeaders([]string{"colA", "colC"})
func (df *Dataframe) RemoveColumnsByHeaders(headers []string) {
	idx := make([]int, len(headers))
	for i, header := range headers {
		idx[i] = fnVisual.IndexOf(header, df.headers)
	}
	df.RemoveColumns(idx)
}

// RemoveLines removes rows at the provided zero-based indices.
// Indices outside valid range are ignored. Order of remaining rows is preserved.
// Implementation detail: it builds a boolean mask and applies it to all columns.
// Examples:
//
//	df.RemoveLines([]int{1, 3})
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

// Compute builds a new Series by applying a row-wise function across the dataframe.
// The provided function fn receives the dataframe and the current row index and
// should return a value compatible with the target type t.
// Params:
// - t: resulting series logical type (e.g., "string", "number", "float", "bool", "date").
// - fn: callback invoked for each row index.
// Returns: a new Series[any] of length equal to the number of rows.
// Examples:
//
//	// Sum two numeric columns into a new number series
//	out := df.Compute("number", func(d *df.Dataframe, i int) any {
//		c1, _ := d.GetSeriesByHeader("A")
//		c2, _ := d.GetSeriesByHeader("B")
//		return c1.GetValue(i).(int) + c2.GetValue(i).(int)
//	})
func (df *Dataframe) Compute(t string, fn func(d *Dataframe, idx int) any) series.Series[any] {
	rows := df.Shape()[0]
	slice := make([]any, rows)
	s := series.New(slice, t)
	for i := 0; i < rows; i++ {
		s = s.SetValue(i, fn(df, i))
	}
	return s
}

// Debug prints the dataframe to stdout in a simple aligned table format.
// It includes headers with their types and all rows. For an empty dataframe,
// it prints a placeholder line.
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
