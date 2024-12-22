package day22

import (
	"adventofcode2024/pkg/utils"
)

const ITERATIONS = 2000
const PRUNE = 16777216

func PartOne(path string) int {
	return utils.Reduce(parse(path), 0, func(accum int, secret int) int {
		for range ITERATIONS {
			secret = nextSecret(secret)
		}
		return accum + secret
	})
}

func PartTwo(path string) int {
	var deltas [4]int
	totals := make(map[[4]int]int)

	for _, initial := range parse(path) {
		previous := initial
		purchases := make(map[[4]int]int)

		for idx := range ITERATIONS {
			next := nextSecret(previous)
			delta := (next % 10) - (previous % 10)
			previous = next

			deltas[0], deltas[1], deltas[2] = deltas[1], deltas[2], deltas[3]
			deltas[3] = delta

			if _, exists := purchases[deltas]; !exists && idx > 4 {
				purchases[deltas] = next % 10
			}
		}

		for key, value := range purchases {
			totals[key] += value
		}
	}

	return utils.MaxMapValue(totals)
}

func nextSecret(number int) int {
	number = (number ^ (number * 64)) % PRUNE
	number = (number ^ (number / 32)) % PRUNE
	number = (number ^ (number * 2048)) % PRUNE
	return number
}

func parse(path string) []int {
	return utils.Map(utils.MustReadInput(path), func(line string) int {
		return utils.MustAtoi(line)
	})
}
