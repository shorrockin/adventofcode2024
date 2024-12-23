package day02

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/slices"
	"strings"
)

func Solve(path string, allowSkip bool) int {
	reports := slices.Map(utils.MustReadInput(path), func(line string) []int {
		parts := strings.Fields(line)
		return slices.Map(parts, utils.MustAtoi)
	})

	return slices.Count(reports, func(report []int) bool {
		if allowSkip {
			return isSafeWithSkip(report)
		}
		return isSafe(report)
	})

}

func isSafe(input []int) bool {
	descending := input[0] > input[1]
	for idx := 1; idx < len(input); idx++ {
		delta := 0
		if descending {
			delta = input[idx-1] - input[idx]
		} else {
			delta = input[idx] - input[idx-1]
		}
		if delta > 3 || delta <= 0 {
			return false
		}

	}
	return true
}

func isSafeWithSkip(report []int) bool {
	if isSafe(report) {
		return true
	}

	for skip := range len(report) {
		var skipped []int = make([]int, len(report))
		copy(skipped, report)
		skipped = append(skipped[:skip], skipped[skip+1:]...)

		if isSafe(skipped) {
			return true
		}
	}

	return false
}
