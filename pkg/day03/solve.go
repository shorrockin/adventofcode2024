package day03

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/assert"
	"regexp"
	"strings"
)

type Part int

const (
	PartOne Part = iota
	PartTwo
)

func Solve(path string, part Part) int {
	pattern := regexp.MustCompile(`(mul\(\d+,\d+\)|don't\(\)|do\(\))`)
	product := 0
	skipping := false

	for _, line := range utils.MustReadInput(path) {
		for _, match := range pattern.FindAllString(line, -1) {
			switch match {
			case "do()":
				skipping = false
			case "don't()":
				skipping = true
			default:
				if skipping && part == PartTwo {
					continue
				}

				parts := strings.Split(match, ",")
				assert.Equal(2, len(parts), "expected 2 parts in mul(x,y) expression", parts)
				product += utils.MustAtoi(parts[0][4:]) * utils.MustAtoi(parts[1][:len(parts[1])-1])
			}
		}
	}

	return product
}
