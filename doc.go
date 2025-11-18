// Package main documents the go-starter project itself.
//
// Overview
//
//	go-starter is a small, generic data toolkit for Go. It centers around two
//	core concepts:
//	- series: a typed, column-like slice wrapper with rich operations
//	- df (dataframe): a minimal 2D table built from multiple series
//
//	Around these you get:
//	- extract: helpers to load data (CSV/Excel) into series/dataframes
//	- fn: functional helpers (Map, Filter, Reduce, Any/All, IndexOf, Reverse)
//	- is: predicates and small utilities (In, Zero, Truthy/Falsy, SameSlice)
//
//	Import and use the subpackages directly in your code. This top-level
//	package exists only to host these project-level docs (pkg.go.dev root page)
//	and possibly minimal examples. There is no public API in package main.
//
// Installation
//
//	go get github.com/visual-pivert/go-starter@latest
//
//	Requires Go 1.24 or newer (see go.mod).
//
// Quickstart
//
//	Build a dataframe from two series and do simple operations.
//
//	package yourpkg
//
//	import (
//	    "fmt"
//	    "github.com/visual-pivert/go-starter/df"
//	    "github.com/visual-pivert/go-starter/series"
//	)
//
//	func exampleQuickstart() {
//	    frame := df.New(
//	        []series.Series[any]{
//	            // Note: df.Dataframe uses Series[any] columns.
//	            series.New([]any{1, 2, 3, 4}, "number"),
//	            series.New([]any{"a", "b", "c", "d"}, "string"),
//	        },
//	        []string{"nums", "names"},
//	    )
//
//	    // Shape: [rows, cols]
//	    shape := frame.Shape()
//	    fmt.Println("shape:", shape)
//
//	    // Select a column by header
//	    col, header := frame.GetSeriesByHeader("nums")
//	    fmt.Println("header:", header, "len:", col.Len())
//	}
//
// Core concepts
//
//	Series[T]
//	- Construct with series.New([]T, tTag), where tTag is one of:
//	  "string", "number", "float", "bool", "date".
//	- Provides operations like Map, Filter, Reduce (see tests for more usage).
//	- When building dataframes, use Series[any] columns.
//
//	Dataframe
//	- Construct with df.New([]series.Series[any], headers).
//	- Retrieve columns by index or by header: GetSeries(i) or GetSeriesByHeader(name).
//	- Filter rows with a boolean mask: ApplyFromBoolStatement(mask).
//	- Copy creates a shallow copy that shares underlying series.
//
// Data ingestion (extract)
//
//	CSV example (see extract/csv.go):
//
//	  package yourpkg
//
//	  import (
//	      "github.com/visual-pivert/go-starter/extract"
//	      "github.com/visual-pivert/go-starter/series"
//	  )
//
//	  func loadCSVExample(path string) ([]series.Series[any], error) {
//	      // Refer to extract package for concrete functions and options.
//	      // Typical flow: extract CSV rows, then convert to series/dataframe.
//	      // See tests for concrete usage examples.
//	      return nil, nil
//	  }
//
// Functional helpers (fn)
//
//	The fn package gives small utilities for slice/series operations:
//	- Map, Filter, Reduce, Any, All, IndexOf, Reverse
//	See fn/*.go and fn/*_test.go for idiomatic patterns.
//
// Predicates/utilities (is)
//
//	The is package provides small helpers like:
//	- In, Zero, Truthy/Falsy, SameSlice
//	See is/*.go and tests for usage.
//
// Recipes
//
//	Build a boolean mask and filter a dataframe:
//
//	  package yourpkg
//
//	  import (
//	      "github.com/visual-pivert/go-starter/df"
//	      "github.com/visual-pivert/go-starter/series"
//	  )
//
//	  func filterRowsExample() *df.Dataframe {
//	      ages := series.New([]any{16, 21, 34, 15, 18}, "number")
//	      names := series.New([]any{"Ana", "Bob", "Cyd", "Dan", "Eli"}, "string")
//	      frame := df.New([]series.Series[any]{ages, names}, []string{"age", "name"})
//
//	      // Build boolean mask: age >= 18. The boolean series must match row count.
//	      keep := series.New([]bool{false, true, true, false, true}, "bool")
//
//	      frame.ApplyFromBoolStatement(keep)
//	      return frame
//	  }
//
// Where to look next
//
//   - df package docs:      github.com/visual-pivert/go-starter/df
//
//   - series package docs:  github.com/visual-pivert/go-starter/series
//
//   - extract package docs: github.com/visual-pivert/go-starter/extract
//
//   - fn package docs:      github.com/visual-pivert/go-starter/fn
//
//   - is package docs:      github.com/visual-pivert/go-starter/is
//
//     The *_test.go files across packages are a great source of working examples.
//     On pkg.go.dev, navigate to the subpackages above to see full APIs.
//
// Versioning and stability
//
//	This project is small and evolving. Pin a version/tag when using it in
//	production code. If you need a new helper or find a bug, please open an
//	issue or PR.
package main
