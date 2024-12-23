package day01

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/assert"
	. "adventofcode2024/pkg/utils/slices"
	"math"
	"slices"
	"strings"
)

type Part int

const (
	PartOne Part = iota
	PartTwo
)

func Solve(path string, part Part) int {
	var left []int
	var right []int

	for _, line := range utils.MustReadInput(path) {
		ints := strings.Fields(line)
		assert.Equal(2, len(ints), "expected 2 fields per line", len(ints))
		left = append(left, utils.MustAtoi(ints[0]))
		right = append(right, utils.MustAtoi(ints[1]))
	}

	if part == PartOne {
		slices.Sort(left)
		slices.Sort(right)

		sum := 0
		for idx, l := range left {
			sum += int(math.Abs(float64(right[idx] - l)))
		}
		return sum
	} else {
		frequencies := make(map[int]int)
		for _, r := range right {
			frequencies[r]++
		}

		return Reduce(left, 0, func(acc int, l int) int {
			return acc + (l * frequencies[l])
		})
	}
}
