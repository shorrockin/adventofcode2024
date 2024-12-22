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
	secrets := make(map[int][]int)

	for _, initial := range parse(path) {
		secrets[initial] = make([]int, ITERATIONS)
		previous := initial

		for idx := range ITERATIONS {
			next := nextSecret(previous)
			secrets[initial][idx] = next % 10
			previous = next
		}
	}

	total := make(map[[4]int]int)
	for initial, digits := range secrets {
		purchases := make(map[[4]int]int)

		for idx := 0; idx+4 < len(digits); idx++ {
			key := [4]int{}

			previous := initial
			if idx > 0 {
				previous = digits[idx-1]
			}

			for offset := range 4 {
				value := digits[idx+offset] - previous
				key[offset] = value
				previous = digits[idx+offset]

			}

			// they'll purchase the first time they see this sequence,
			// so only update the map if it's not there
			if _, exists := purchases[key]; !exists {
				purchases[key] = previous
			}
		}

		for key, value := range purchases {
			total[key] += value
		}
	}

	return utils.MaxMapValue(total)
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
