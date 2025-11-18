package extract

import (
	"archive/zip"
	"encoding/xml"
	"io"
	pathpkg "path"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/visual-pivert/go-starter/df"
)

type sst struct {
	SI []si `xml:"si"`
}

type workbook struct {
	Sheets     []wbSheet `xml:"sheets>sheet"`
	WorkbookPr wbPr      `xml:"workbookPr"`
}

type wbPr struct {
	Date1904 bool `xml:"date1904,attr"`
}

type wbSheet struct {
	Name string `xml:"name,attr"`
	RID  string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
}

type relationships struct {
	Items []rel `xml:"Relationship"`
}

type rel struct {
	ID     string `xml:"Id,attr"`
	Target string `xml:"Target,attr"`
}

type styles struct {
	NumFmts numFmts `xml:"numFmts"`
	CellXfs cellXfs `xml:"cellXfs"`
}

type numFmts struct {
	Fmts []numFmt `xml:"numFmt"`
}

type numFmt struct {
	ID   int    `xml:"numFmtId,attr"`
	Code string `xml:"formatCode,attr"`
}

type cellXfs struct {
	Xfs []xf `xml:"xf"`
}

type xf struct {
	NumFmtId int `xml:"numFmtId,attr"`
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
	S  int       `xml:"s,attr"` // style index
	V  string    `xml:"v"`
	IS inlineStr `xml:"is"`
}

type inlineStr struct {
	T string `xml:"t"`
}

// Excel reads an .xlsx file and returns a Dataframe (check out df package for more info).
// Examples:
//
// extract.Excel("data.xlsx", "Sheet1", []string{"int", "string"}, 0) // return dataframe
func Excel(path string, sheet string, types []string, headerIdx int) *df.Dataframe {
	reader, err := zip.OpenReader(path)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	// Load workbook to resolve sheet name and date1904 setting
	var wb workbook
	date1904 := false
	if f := findInZip(&reader.Reader, "xl/workbook.xml"); f != nil {
		rc, err := f.Open()
		if err == nil {
			data, _ := io.ReadAll(rc)
			_ = rc.Close()
			_ = xml.Unmarshal(data, &wb)
			date1904 = wb.WorkbookPr.Date1904
		}
	}

	// Map r:id -> target path for sheets
	relMap := map[string]string{}
	if f := findInZip(&reader.Reader, "xl/_rels/workbook.xml.rels"); f != nil {
		rc, err := f.Open()
		if err == nil {
			data, _ := io.ReadAll(rc)
			_ = rc.Close()
			var rels relationships
			if xml.Unmarshal(data, &rels) == nil {
				for _, it := range rels.Items {
					relMap[it.ID] = it.Target
				}
			}
		}
	}

	// Resolve sheet target by human-readable name first
	var sheetPath string
	for _, s := range wb.Sheets {
		if s.Name == sheet {
			if tgt, ok := relMap[s.RID]; ok {
				// Normalize target to avoid leading slash dropping the 'xl' prefix
				if strings.HasPrefix(tgt, "/") {
					tgt = strings.TrimPrefix(tgt, "/")
				}
				sheetPath = pathpkg.Join("xl", tgt)
				break
			}
		}
	}

	// Fallbacks: if not found by name, try legacy behavior
	var sheetFile *zip.File
	if sheetPath != "" {
		sheetFile = findInZip(&reader.Reader, sheetPath)
	}
	if sheetFile == nil {
		// Accept raw xml filename or suffix like "sheet1.xml"
		target := sheet
		if !strings.HasSuffix(target, ".xml") {
			target = sheet + ".xml"
		}
		for _, file := range reader.File {
			if strings.HasSuffix(file.Name, target) {
				sheetFile = file
				break
			}
		}
	}
	if sheetFile == nil {
		panic("sheet not found in xlsx: " + sheet)
	}

	// Load shared strings
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

	// Load styles for date/time detection
	styleNumFmt := []int{}
	customFmt := map[int]string{}
	if f := findInZip(&reader.Reader, "xl/styles.xml"); f != nil {
		rc, err := f.Open()
		if err == nil {
			data, _ := io.ReadAll(rc)
			_ = rc.Close()
			var st styles
			if xml.Unmarshal(data, &st) == nil {
				for _, nf := range st.NumFmts.Fmts {
					customFmt[nf.ID] = nf.Code
				}
				for _, x := range st.CellXfs.Xfs {
					styleNumFmt = append(styleNumFmt, x.NumFmtId)
				}
			}
		}
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
			rowsMap[r.R][col] = cellValue(c, shared, styleNumFmt, customFmt, date1904)
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

	return df.FromRaw(out, types, headerIdx)
}

func findInZip(r *zip.Reader, name string) *zip.File {
	for _, f := range r.File {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func cellValue(c cell, shared []string, styleNumFmt []int, customFmt map[int]string, date1904 bool) string {
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
		// Possible number/date/plain
		v := strings.TrimSpace(c.V)
		if v == "" {
			return ""
		}
		// Check style for date/time
		if c.S >= 0 && c.S < len(styleNumFmt) {
			numFmtId := styleNumFmt[c.S]
			fmtCode, hasCustom := customFmt[numFmtId]
			if isDateNumFmt(numFmtId) || (hasCustom && isDateFormatCode(fmtCode)) {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					return excelSerialToISOString(f, date1904)
				}
			}
		}
		return v
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

// Determine if a built-in numFmtId represents a date/time in Excel
func isDateNumFmt(id int) bool {
	// Built-in date/time codes per Excel spec
	if (id >= 14 && id <= 22) || (id >= 27 && id <= 36) || (id >= 45 && id <= 47) {
		return true
	}
	return false
}

// Heuristic check whether a custom format code is date/time-like
func isDateFormatCode(code string) bool {
	lc := strings.ToLower(code)
	// Remove escaped sections and literals in quotes to reduce false positives
	// but keep a simple heuristic: if it contains any y/m/d/h/s or brackets like [h]
	if strings.ContainsAny(lc, "ymdhs") {
		return true
	}
	if strings.Contains(lc, "[h]") || strings.Contains(lc, "[mm]") || strings.Contains(lc, "[ss]") {
		return true
	}
	return false
}

// Convert Excel serial date (with fractional time) to ISO 8601 string using 1900/1904 systems
func excelSerialToISOString(serial float64, date1904 bool) string {
	tm := excelSerialToTime(serial, date1904)
	return tm.Format("2006-01-02T15:04:05")
}

// Convert Excel serial to time.Time in UTC
func excelSerialToTime(serial float64, date1904 bool) time.Time {
	// Separate days and fractional day
	days := int(serial)
	frac := serial - float64(days)
	var base time.Time
	if date1904 {
		// In the 1904 date system, serial 0 corresponds to 1904-01-01
		base = time.Date(1904, 1, 1, 0, 0, 0, 0, time.UTC)
	} else {
		// In the 1900 system, Excel incorrectly treats 1900 as a leap year.
		// A common approach is to use base 1899-12-30 and subtract one day for serials >= 60
		base = time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
		if days >= 60 { // account for the non-existent 1900-02-29
			days--
		}
	}
	sec := int64(frac*86400.0 + 0.5) // round to nearest second
	return base.AddDate(0, 0, days).Add(time.Duration(sec) * time.Second)
}
