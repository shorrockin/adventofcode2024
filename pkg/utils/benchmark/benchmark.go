package benchmark

import (
	"fmt"
	"time"
)

var names = make(map[string]int)

func Profile(name string, fn func()) {
	ProfileAndReturn(name, func() bool {
		fn()
		return true
	})
}

func ProfileAndReturn[T any](name string, fn func() T) T {
	start := time.Now()
	names[name]++
	if names[name] != 1 {
		name = fmt.Sprintf("%s-%d", name, names[name])
	}

	fmt.Println()
	fmt.Printf("[%s] starting at %s\n", name, start.Format(time.RFC3339))
	value := fn()
	elapsed := time.Since(start)
	fmt.Printf("[%s] finished, took %s\n", name, elapsed)
	return value
}
