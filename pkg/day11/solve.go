package day11

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/slices"
	"math"
	"strings"
)

type CacheKey struct {
	value     int
	iteration int
}

var cache = make(map[CacheKey]int)

func Solve(raw string, times int) int {
	return slices.Reduce(parse(raw), 0, func(acc, value int) int {
		return acc + produce(value, times)
	})
}

func produce(value int, iteration int) int {
	if iteration == 0 {
		return 1
	}

	if key, ok := cache[CacheKey{value, iteration}]; ok {
		return key
	}

	if value == 0 {
		res := produce(1, iteration-1)
		cache[CacheKey{value, iteration}] = res
		return res
	}

	left, right, even := split(value)
	if even {
		res := produce(left, iteration-1) + produce(right, iteration-1)
		cache[CacheKey{value, iteration}] = res
		return res
	}

	res := produce(value*2024, iteration-1)
	cache[CacheKey{value, iteration}] = res
	return res
}

func split(value int) (int, int, bool) {
	length := int(math.Log10(float64(value))) + 1
	if length%2 != 0 {
		return 0, 0, false
	}

	middle := int(math.Pow10(length / 2))
	left := value / middle
	right := value % middle

	return left, right, true
}

func parse(raw string) []int {
	return slices.Map(strings.Fields(raw), func(val string) int {
		return utils.MustAtoi(val)
	})
}
