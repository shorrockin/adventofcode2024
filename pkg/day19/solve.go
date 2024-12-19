package day19

import (
	"adventofcode2024/pkg/utils"
	"strings"
)

func Solve(path string, partOne bool) int {
	patterns, designs := parse(path)
	cache := make(map[string]int)

	return utils.Reduce(designs, 0, func(acc int, design string) int {
		count := CountPatterns(design, &patterns, cache)
		if partOne && count > 0 {
			count = 1
		}
		return acc + count
	})
}

func CountPatterns(design string, patterns *[]string, cache map[string]int) int {
	if count, ok := cache[design]; ok {
		return count
	}

	if len(design) == 0 {
		return 1
	}

	count := 0
	for _, pattern := range *patterns {
		if strings.HasPrefix(design, pattern) {
			count += CountPatterns(design[len(pattern):], patterns, cache)
		}
	}

	cache[design] = count
	return count
}

func parse(path string) ([]string, []string) {
	lines := utils.MustReadInput(path)
	patterns := strings.Split(lines[0], ", ")
	designs := make([]string, 0)

	for _, line := range lines[2:] {
		designs = append(designs, line)
	}

	return patterns, designs
}
