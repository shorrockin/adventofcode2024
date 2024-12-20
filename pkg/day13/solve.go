package day13

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/assert"
	"math"
	"regexp"
	"strings"
)

type Point struct {
	x, y float64
}

type ClawMachine struct {
	buttonA Point
	buttonB Point
	target  Point
}

func Solve(path string, partOne bool) int {
	return utils.Reduce(parse(path, partOne), 0, func(acc int, machine ClawMachine) int {
		if x, y, ok := solveSystem(machine); ok {
			if !partOne || (x <= 100 && y <= 100) {
				return acc + (x * 3) + y
			}
		}
		return acc
	})
}

func solveSystem(machine ClawMachine) (int, int, bool) {
	determinant := (machine.buttonA.x * machine.buttonB.y) - (machine.buttonA.y * machine.buttonB.x)
	if determinant == 0 {
		return 0, 0, false
	}

	// use cramer's rule to solve for x and y
	x := ((machine.target.x * machine.buttonB.y) - (machine.target.y * machine.buttonB.x)) / determinant
	y := ((machine.buttonA.x * machine.target.y) - (machine.target.x * machine.buttonA.y)) / determinant

	if math.Mod(x, 1) == 0 && math.Mod(y, 1) == 0 {
		return int(x), int(y), true
	}

	return 0, 0, true
}

func parse(path string, partOne bool) []ClawMachine {
	definitions := strings.Split(strings.Join(utils.MustReadInput(path), "\n"), "\n\n")
	buttonRegexp := regexp.MustCompile(`Button [A|B]: X\+(\d+), Y\+(\d+)`)
	prizeRegexp := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	return utils.Map(definitions, func(definition string) ClawMachine {
		lines := strings.Split(definition, "\n")
		assert.Equal(len(lines), 3, "expected 3 lines in definition")

		buttonA := buttonRegexp.FindAllStringSubmatch(lines[0], -1)
		buttonB := buttonRegexp.FindAllStringSubmatch(lines[1], -1)
		target := prizeRegexp.FindAllStringSubmatch(lines[2], -1)
		increase := 0

		if !partOne {
			increase = 10000000000000
		}

		return ClawMachine{
			buttonA: NewPoint(buttonA[0][1], buttonA[0][2], 0),
			buttonB: NewPoint(buttonB[0][1], buttonB[0][2], 0),
			target:  NewPoint(target[0][1], target[0][2], increase),
		}
	})
}

func NewPoint(x, y string, increase int) Point {
	return Point{
		x: float64(utils.MustAtoi(x) + increase),
		y: float64(utils.MustAtoi(y) + increase),
	}
}
