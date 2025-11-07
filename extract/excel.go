package extract

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/visual-pivert/go-starter/df"
)

type sst struct {
	SI []si `xml:"si"`
}

type si struct {
	T  string `xml:"t"`
	Rs []r    `xml:"r"`
}

type r struct {
	T string `xml:"t"`
}

type worksheet struct {
	SheetData sheetData `xml:"sheetData"`
}

type sheetData struct {
	Rows []row `xml:"row"`
}

type row struct {
	R  int    `xml:"r,attr"`
	Cs []cell `xml:"c"`
}

type cell struct {
	R  string    `xml:"r,attr"` // e.g., A1
	T  string    `xml:"t,attr"` // cell type, e.g., s (shared), b (bool), inlineStr
	V  string    `xml:"v"`
	IS inlineStr `xml:"is"`
}

type inlineStr struct {
	T string `xml:"t"`
}

func Excel(path string, sheet string, headerIdx int) *df.Df {
	sheet = sheet + ".xml"
	reader, err := zip.OpenReader(path)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	// Load shared strings if present
	shared := []string{}
	if f := findInZip(&reader.Reader, "xl/sharedStrings.xml"); f != nil {
		rc, err := f.Open()
		if err == nil {
			data, _ := io.ReadAll(rc)
			_ = rc.Close()
			var s sst
			if xml.Unmarshal(data, &s) == nil {
				for _, item := range s.SI {
					if len(item.Rs) > 0 {
						var b strings.Builder
						for _, rr := range item.Rs {
							b.WriteString(rr.T)
						}
						shared = append(shared, b.String())
					} else {
						shared = append(shared, item.T)
					}
				}
			}
		}
	}

	// Find the requested sheet by suffix match, e.g., "sheet1.xml"
	var sheetFile *zip.File
	for _, file := range reader.File {
		if strings.HasSuffix(file.Name, sheet) {
			sheetFile = file
			break
		}
	}
	if sheetFile == nil {
		panic("sheet not found in xlsx: " + sheet)
	}

	rc, err := sheetFile.Open()
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(rc)
	_ = rc.Close()
	if err != nil {
		panic(err)
	}

	var ws worksheet
	if err := xml.Unmarshal(data, &ws); err != nil {
		panic(err)
	}

	// Build a map of row index -> []string by placing cells at their column positions
	rowsMap := map[int]map[int]string{}
	maxCol := 0
	rowIndices := []int{}
	for _, r := range ws.SheetData.Rows {
		if _, ok := rowsMap[r.R]; !ok {
			rowsMap[r.R] = map[int]string{}
			rowIndices = append(rowIndices, r.R)
		}
		for _, c := range r.Cs {
			col := colRefToIndex(c.R)
			if col > maxCol {
				maxCol = col
			}
			rowsMap[r.R][col] = cellValue(c, shared)
		}
	}

	// sort rows by index
	sort.Ints(rowIndices)
	// construct [][]string
	var out [][]string
	width := maxCol + 1
	for _, ri := range rowIndices {
		rowMap := rowsMap[ri]
		line := make([]string, width)
		for i := 0; i < width; i++ {
			if v, ok := rowMap[i]; ok {
				line[i] = v
			} else {
				line[i] = ""
			}
		}
		out = append(out, line)
	}

	return df.FromRaw(out, headerIdx)
}

func findInZip(r *zip.Reader, name string) *zip.File {
	for _, f := range r.File {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func cellValue(c cell, shared []string) string {
	t := c.T
	switch t {
	case "s": // shared string
		idx, err := strconv.Atoi(strings.TrimSpace(c.V))
		if err == nil && idx >= 0 && idx < len(shared) {
			return shared[idx]
		}
		return ""
	case "inlineStr":
		return c.IS.T
	case "b":
		if strings.TrimSpace(c.V) == "1" {
			return "true"
		}
		return "false"
	default:
		// number, date serial, or plain text without type attribute
		return strings.TrimSpace(c.V)
	}
}

// Convert a cell reference like "C5" or "AA10" to zero-based column index.
func colRefToIndex(ref string) int {
	// Extract letters prefix
	lettersEnd := 0
	for lettersEnd < len(ref) {
		ch := ref[lettersEnd]
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			lettersEnd++
			continue
		}
		break
	}
	letters := ref[:lettersEnd]
	letters = strings.ToUpper(letters)
	// Convert base-26 letters to index
	col := 0
	for i := 0; i < len(letters); i++ {
		col = col*26 + int(letters[i]-'A'+1)
	}
	return col - 1
}
