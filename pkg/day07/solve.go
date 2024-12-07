package day07

import (
	"adventofcode2024/pkg/utils"
	"slices"
	"strconv"
	"strings"
)

type Equation struct {
	result int
	input  []int
}

func Solve(path string, partOne bool) int {
	formulas := parse(path)
	sum := 0

	for _, formula := range formulas {
		solutions := []int{formula.input[0]}

		for inputIdx := 1; inputIdx < len(formula.input); inputIdx++ {
			input := formula.input[inputIdx]

			for idx, solution := range solutions {
				if !partOne {
					joined := strconv.Itoa(solution) + strconv.Itoa(input)
					solutions = append(solutions, utils.MustAtoi(joined))
				}
				solutions = append(solutions, solution*input)
				solutions[idx] = solution + input
			}
		}

		if slices.Contains(solutions, formula.result) {
			sum += formula.result
		}

	}

	return sum
}

func parse(path string) []Equation {
	return utils.Map(utils.MustReadInput(path), func(line string) Equation {
		fields := strings.Fields(line)
		return Equation{
			result: utils.MustAtoi(strings.TrimSuffix(fields[0], ":")),
			input: utils.Map(fields[1:], func(value string) int {
				return utils.MustAtoi(value)
			}),
		}
	})
}
