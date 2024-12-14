package day14

import (
	"adventofcode2024/pkg/assert"
	"adventofcode2024/pkg/grid"
	"adventofcode2024/pkg/utils"
	"regexp"
)

type Robot struct {
	position grid.Coordinate
	velocity grid.Coordinate
}

func (r *Robot) Move(width, height int) {
	r.position = r.position.Offset(r.velocity).Bounded(width, height)
}

func Solve(path string, width, height int, partOne bool) int {
	robots := parse(path)

	if partOne {
		return countQuadrants(robots, width, height)
	}

	return treeLocation(robots, width, height)
}

func treeLocation(robots []*Robot, width, height int) int {
	seconds := 0

	for {
		seconds++
		lookup := make(map[grid.Coordinate]*Robot)
		for _, robot := range robots {
			robot.Move(width, height)
			lookup[robot.position] = robot
		}

		// look for a cascading pattern of robots down and to the right
		// of a specific length
		movement := grid.At(1, 1)
		for position := range lookup {
			for iteration := range 6 {
				position = position.Offset(movement)
				if lookup[position] == nil {
					break
				}
				if iteration == 5 {
					display(robots, width, height)
					return seconds
				}

			}
		}
	}
}

func countQuadrants(robots []*Robot, width, height int) int {
	for range 100 {
		for _, robot := range robots {
			robot.Move(width, height)
		}
	}

	quadrants := make([]int, 4)
	for _, robot := range robots {
		if robot.position.X < width/2 && robot.position.Y < height/2 {
			quadrants[0]++
		} else if robot.position.X > width/2 && robot.position.Y < height/2 {
			quadrants[1]++
		} else if robot.position.X < width/2 && robot.position.Y > height/2 {
			quadrants[2]++
		} else if robot.position.X > width/2 && robot.position.Y > height/2 {
			quadrants[3]++
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func display(robots []*Robot, width, height int) {
	lookup := make(map[grid.Coordinate]*Robot)
	for _, robot := range robots {
		lookup[robot.position] = robot
	}

	for y := range height {
		for x := range width {
			if _, ok := lookup[grid.At(x, y)]; ok {
				print("█")
			} else {
				print(" ")
			}
		}
		println()
	}
}

func parse(path string) []*Robot {
	parser := regexp.MustCompile(`p=(.*),(.*) v=(.*),(.*)`)
	return utils.Map(utils.MustReadInput(path), func(line string) *Robot {
		parts := parser.FindStringSubmatch(line)
		assert.Equal(len(parts), 5, "expected 5 parts in line")
		return &Robot{
			position: grid.At(utils.MustAtoi(parts[1]), utils.MustAtoi(parts[2])),
			velocity: grid.At(utils.MustAtoi(parts[3]), utils.MustAtoi(parts[4])),
		}
	})
}
