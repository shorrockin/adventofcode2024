package day17

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/assert"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Computer struct {
	a       int
	b       int
	c       int
	program []int
	output  []int
}

func PartOne(path string) []int {
	computer := parse(path)
	Run(&computer, true)
	return computer.output
}

func PartTwo(path string) int {
	computer := parse(path)

	value := 1
	for {
		computer.reset(value)
		Run(&computer, false)

		if len(computer.program) == len(computer.output) {
			if utils.EqualSlices(computer.program, computer.output) {
				fmt.Printf("found: %v\n", value)
				return value
			}

			for idx := len(computer.program) - 1; idx >= 0; idx-- {
				if computer.program[idx] != computer.output[idx] {
					// each digit is determined by 3 bits in reverse order,
					// that means that the first value is the output is
					// determined by the last 3 bits, while the last value is
					// deterimned by the first 3 bits.
					//
					// this means that we can change the first value in the
					// output (idx 1) by incrementing by 1, and interestingly,
					// the second value (idx 2) can be changed by incrementing
					// by 8, then 64, then 512, etc. more succinctly
					// incrementing by 8^0, 8^1, 8^2, controls the output index
					// we are changing, where the power is index
					//
					// with this if we match against last value in the output
					// (idx 15), we will then increment by 8^15 until we have
					// the correct value at this index. once we do, it
					// stabilizes, and we can work on the next value, 8^14, etc
					// until we have a solution.
					previous := value
					value += int(math.Pow(8, float64(idx)))

					fmt.Printf("        matches: %v\n", len(computer.program)-idx-1)
					fmt.Printf("         output: %+v\n", computer.output)
					fmt.Printf("          value: %v\n", previous)
					fmt.Printf("           next: %v\n", value)
					fmt.Printf("      increment: 8^%d (%v)\n", idx, math.Pow(8, float64(idx)))
					fmt.Printf("binary previous: %+v\n", strconv.FormatInt(int64(previous), 2))
					fmt.Printf(" binary current: %+v\n", strconv.FormatInt(int64(value), 2))
					fmt.Println()

					break
				}
			}
		} else if len(computer.program) > len(computer.output) {
			value *= 2
		}
	}
}

func (c *Computer) reset(a int) {
	c.a = a
	c.b = 0
	c.c = 0
	c.output = c.output[:0]
}

func Run(computer *Computer, partOne bool) {
	for instruction := 0; instruction < len(computer.program); instruction = instruction + 2 {
		opcode := computer.program[instruction]
		literal := computer.program[instruction+1]
		combo := literal

		switch combo {
		case 4:
			combo = computer.a
		case 5:
			combo = computer.b
		case 6:
			combo = computer.c
		default:
			if combo > 6 {
				panic(assert.Fail("combo operand cannot be greater than 6", "value", combo))
			}
		}

		switch opcode {
		case 0:
			computer.a = computer.a / int(math.Pow(2, float64(combo)))
		case 1:
			computer.b = computer.b ^ literal
		case 2:
			computer.b = combo % 8
		case 3:
			if computer.a != 0 {
				instruction = literal - 2
			}
		case 4:
			computer.b = computer.b ^ computer.c
		case 5:
			computer.output = append(computer.output, combo%8)
		case 6:
			computer.b = computer.a / int(math.Pow(2, float64(combo)))
		case 7:
			computer.c = computer.a / int(math.Pow(2, float64(combo)))
		default:
			panic(assert.Fail("unknown opcode instruction", "value", opcode))
		}
	}
}

func parse(path string) Computer {
	lines := utils.MustReadInput(path)

	return Computer{
		a: utils.MustAtoi(lines[0][12:]),
		b: utils.MustAtoi(lines[1][12:]),
		c: utils.MustAtoi(lines[2][12:]),
		program: utils.Map(strings.Split(lines[4][9:], ","), func(part string) int {
			return utils.MustAtoi(part)
		}),
	}
}
