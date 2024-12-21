package benchmark

import (
	"fmt"
	"time"
)

type Benchmark struct {
	start time.Time
	last  time.Time
	name  string
	laps  int
}

var names = make(map[string]int)

func Run(name string, fn func()) {
	AndReturn(name, func() bool {
		fn()
		return true
	})
}

func AndReturn[T any](name string, fn func() T) T {
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

func Start(name string) Benchmark {
	names[name]++
	if names[name] != 1 {
		name = fmt.Sprintf("%s-%d", name, names[name])
	}
	start := time.Now()
	return Benchmark{start, start, name, 0}
}

func (b *Benchmark) Measure(name string) {
	elapsed := time.Since(b.last)
	total := time.Since(b.start)

	b.laps++
	fmt.Printf("[%s][%s][%d] took %s (total: %s)\n", b.name, name, b.laps, elapsed, total)
	b.last = time.Now()
}
