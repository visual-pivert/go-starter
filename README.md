# go-starter

Lightweight, practical building blocks for working with slices, series, and simple dataframes in Go. Comes with:
- Generic functional helpers (`fn`): Map, Reduce, Filter, Reverse, Any/All, IndexOf, etc.
- A small but handy `series` type with chainable methods.
- A minimal `df` (dataframe) structure for columnar data.
- `extract` helpers to load CSV (and basic Excel parsing planned).
- `is` helpers for truthiness checks.

Important note — early stage/alpha: This project is intentionally tiny and experimental. The code favors simplicity over bullet‑proof error handling. Some functions may panic on invalid inputs or I/O errors (e.g., reading a missing file). Treat this as a learning/utility library, not production‑hardened code yet.

## Installation

```bash
go get github.com/visual-pivert/go-starter
```

Go modules will fetch the packages. Import what you need:

```go
import (
    "github.com/visual-pivert/go-starter/fn"
    "github.com/visual-pivert/go-starter/series"
    "github.com/visual-pivert/go-starter/df"
    "github.com/visual-pivert/go-starter/extract"
    "github.com/visual-pivert/go-starter/is"
)
```

## Quickstart

### Functional helpers (`fn`)
A small collection of generic helpers that work with slices.

```go
package main

import (
    "fmt"
    "github.com/visual-pivert/go-starter/fn"
)

func main() {
    nums := []int{1, 2, 3, 4}

    // Map: fn expects (value, index)
    doubled := fn.Map(nums, func(v int, _ int) int { return v * 2 })
    fmt.Println(doubled) // [2 4 6 8]

    // Filter: keep only even numbers
    evens := fn.Filter(nums, func(v int) bool { return v%2 == 0 })
    fmt.Println(evens) // [2 4]

    // Reduce: cumulative results (cum, value, index)
    // For [1 2 3 4] with +, it returns [1 3 6 10]. Last item is the total.
    cum := fn.Reduce(nums, 0, func(cum int, value int, _ int) int { return cum + value })
    fmt.Println(cum)            // [1 3 6 10]
    fmt.Println(cum[len(cum)-1]) // 10
}
```

Other goodies: `fn.MapReverse`, `fn.FilterI` (indices), `fn.FilterTruthy`, `fn.FilterITruthy`, `fn.FilterToBoolStatement`, `fn.Reverse`, `fn.IndexOf`, `fn.Any`, `fn.All`.

### Series (`series`)
A generic, typed one‑dimensional container with convenience methods.

```go
package main

import (
    "fmt"
    "github.com/visual-pivert/go-starter/series"
)

func main() {
    s := series.New([]int{1, 2, 3, 4}, "number")

    fmt.Println(s.Len(), s.Count())   // 4 4
    fmt.Println(s.ToSlice())          // [1 2 3 4]
    fmt.Println(s.GetValue(2))        // 3

    s2 := s.Map(func(v int, i int) int { return v * (i + 1) })
    fmt.Println(s2.ToSlice())         // [1 4 9 16]

    onlyEven := s.Filter(func(v int) bool { return v%2 == 0 })
    fmt.Println(onlyEven.ToSlice())   // [2 4]

    // Reduce returns a Series with cumulative values
    cum := s.Reduce(0, func(last int, curr int, _ int) int { return last + curr })
    fmt.Println(cum.ToSlice())        // [1 3 6 10]
}
```

Selected methods: `Append`, `AppendTo`, `Pop`, `Shift`, `Remove`, `Range`, `Len`, `Count`, `Type`, `ToSlice`, `Filter`, `FilterI`, `Reduce`, `Map`, `MapToBool`, `ApplyBoolStatement`, `ApplyOrderStatement`, `CountValue`, `GetValue`, `SetValue`, `Reverse`, `Agg`, `Any`, `All`, `IndexOf`.

### Dataframe (`df`)
A minimal column‑oriented structure that composes `series.Series[any]` columns and headers.

```go
package main

import (
    "fmt"
    "github.com/visual-pivert/go-starter/df"
    "github.com/visual-pivert/go-starter/series"
)

func main() {
    d := df.New(nil, []string{})
    d.Append(series.New[any]([]any{"Alice", "Bob"}, "string"), "name")
    d.Append(series.New[any]([]any{23, 31}, "number"), "age")

    fmt.Println(d.Shape()) // [2 2] -> rows, cols

    ageSeries, header := d.GetSeriesByHeader("age")
    fmt.Println(header, ageSeries.ToSlice()) // age [23 31]
}
```

Dataframe ops: `Append`, `Copy`, `Shape`, `GetSeries`, `GetSeriesByHeader`, `RemoveColumns`, `RemoveColumnsByHeaders`, `RemoveLines`, `ApplyFromBoolStatement`, `ApplyFromOrderStatement`, `Compute`, `Debug`.

### Extract (`extract`)
Load data into dataframes.

```go
package main

import (
    "fmt"
    "github.com/visual-pivert/go-starter/extract"
)

func main() {
    // Csv(path, sep, headerIdx, types)
    df := extract.Csv("student_exam_scores.csv", ",", 0, []string{"string", "number"})
    fmt.Println(df.Shape())
}
```

- CSV is supported today via `extract.Csv`.
- Excel helpers exist but are basic/experimental; APIs may change.

### Truthiness helpers (`is`)
Utilities to check for truthy/falsy/zero values across types.

```go
package main

import (
    "fmt"
    "github.com/visual-pivert/go-starter/is"
)

func main() {
    fmt.Println(is.Truthy(1))      // true
    fmt.Println(is.Falsy(""))     // true
    fmt.Println(is.Zero(0))        // true
}
```

## Design goals
- Small, readable implementations you can learn from and copy.
- Generic where it helps, concrete where it keeps things simple.
- Opt‑in ergonomics over heavy abstractions.

## Warnings and stability (please read)
- Early alpha — interfaces may change.
- Some functions panic instead of returning errors (e.g., file reads). Validate inputs and wrap calls if you need safety.
- Not tuned for performance on huge datasets for now.

## Roadmap
- Harden error handling (minimize panics).
- Expand `extract` (better CSV quoting, robust Excel, streaming options).
- More dataframe transforms (grouping, joins, typed schemas).
- Benchmarking/perf passes and docs.

## FAQ
- Why not rely on big data libraries? This project aims to stay tiny, generic, and idiomatic, ideal for small tasks and examples.
- Does `Reduce` return a single value? Here it returns the cumulative slice; take the last value when you want the total.
- Will APIs change? Yes — until a stable v1. Feedback welcome.

## Contributing
PRs and issues are welcome! Please keep changes focused and small. Tests, examples, and doc improvements are especially appreciated.

## License
MIT — see [LICENSE](./LICENSE)
