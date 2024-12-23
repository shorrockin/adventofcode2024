package day22

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/collections"
	"adventofcode2024/pkg/utils/maps"
	"adventofcode2024/pkg/utils/slices"
)

type Sequence [4]int

const ITERATIONS = 2000
const PRUNE = 16777216

func PartOne(path string) int {
	return slices.Reduce(parse(path), 0, func(accum int, secret int) int {
		for range ITERATIONS {
			secret = nextSecret(secret)
		}
		return accum + secret
	})
}

func PartTwo(path string) int {
	totals := make(map[Sequence]int)

	for _, previous := range parse(path) {
		sequence := [4]int{0, 0, 0, 0}
		purchases := collections.NewSet[Sequence]()

		for idx := range ITERATIONS {
			next := nextSecret(previous)
			sequence[0], sequence[1], sequence[2], sequence[3] = sequence[1], sequence[2], sequence[3], (next%10)-(previous%10)
			previous = next

			if idx >= 4 && purchases.MaybeAdd(sequence) {
				totals[sequence] += next % 10
			}
		}
	}

	return maps.MaxValue(totals)
}

func nextSecret(number int) int {
	number = (number ^ (number * 64)) % PRUNE
	number = (number ^ (number / 32)) % PRUNE
	number = (number ^ (number * 2048)) % PRUNE
	return number
}

func parse(path string) []int {
	return slices.Map(utils.MustReadInput(path), func(line string) int {
		return utils.MustAtoi(line)
	})
}
