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
	equations := parse(path)
	sum := 0

	for _, equation := range equations {
		solutions := []int{equation.input[0]}

		for iIdx := 1; iIdx < len(equation.input); iIdx++ {
			input := equation.input[iIdx]

			for sIdx, solution := range solutions {
				if !partOne {
					joined := strconv.Itoa(solution) + strconv.Itoa(input)
					solutions = append(solutions, utils.MustAtoi(joined))
				}
				solutions = append(solutions, solution*input)
				solutions[sIdx] = solution + input
			}
		}

		if slices.Contains(solutions, equation.result) {
			sum += equation.result
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
