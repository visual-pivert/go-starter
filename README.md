# go-starter

Starter utilities and functional helpers for Go — including generic functions like `Map`, `Reduce`, `Filter`, and more.

## Installation

```bash
go get github.com/visual-pivert/go-starter
```

Then import it:

```go
import "github.com/visual-pivert/go-starter/fn"
```

## Examples

```go
package main

import (
	"fmt"

	"github.com/visual-pivert/go-starter/fn"
)

func main() {
	nums := []int{1, 2, 3, 4}

	// Map
	doubled := fn.Map(nums, func(n int) int { return n * 2 })
	fmt.Println(doubled) // [2 4 6 8]

	// Filter
	evens := fn.Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens) // [2 4]

	// Reduce
	sum := fn.Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println(sum)             // [1 3 6 10]
	fmt.Println(sum[len(sum)-1]) // 10
}
```
---

Simple, fast, and reusable — that’s **go-starter**.

## License

This project is licensed under the MIT License — see the [LICENSE](https://github.com/visual-pivert/go-starter/blob/main/LICENSE)
