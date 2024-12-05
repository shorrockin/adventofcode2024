package day05

import (
	"adventofcode2024/pkg/assert"
	"adventofcode2024/pkg/utils"
	"slices"
	"strings"
)

type Part int

const (
	PartOne Part = iota
	PartTwo
)

func Solve(path string, part Part) int {
	rules, updates := parse(path)

	return utils.Reduce(updates, 0, func(accum int, update []string) int {
		for i := 0; i < len(update)-1; i++ {
			left, right := update[i], update[i+1]

			if slices.Contains(rules[right], left) {
				if part == PartOne {
					return accum
				}

				slices.SortFunc(update, func(left, right string) int {
					return slices.Index(rules[left], right)
				})
				return accum + utils.MustAtoi(update[len(update)/2])
			}
		}

		if part == PartOne {
			return accum + utils.MustAtoi(update[len(update)/2])
		}

		return accum
	})
}

func parse(path string) (map[string][]string, [][]string) {
	updates := make([][]string, 0)
	rules := make(map[string][]string)

	for _, line := range utils.MustReadInput(path) {
		if parts := strings.Split(line, "|"); len(parts) == 2 {
			before, after := parts[0], parts[1]
			rules[before] = append(rules[before], after)

		} else if pages := strings.Split(line, ","); len(pages) > 1 {
			assert.Equal(1, len(pages)%2, "expected updates to contain odd number")
			updates = append(updates, pages)
		}
	}
	return rules, updates
}
