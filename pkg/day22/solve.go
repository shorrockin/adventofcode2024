package day22

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/collections"
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

	for _, previous := range parse(path) {
		purchases := collections.NewSet[[4]int]()

		for idx := range ITERATIONS {
			next := nextSecret(previous)
			delta := (next % 10) - (previous % 10)
			deltas[0], deltas[1], deltas[2], deltas[3] = deltas[1], deltas[2], deltas[3], delta
			previous = next

			if !purchases.Contains(deltas) && idx > 4 {
				purchases.Add(deltas)
				totals[deltas] += next % 10
			}
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
