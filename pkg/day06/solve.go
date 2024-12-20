package day06

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/assert"
	"adventofcode2024/pkg/utils/grid"
	"slices"
)

type Path struct {
	coord  grid.Coord
	facing grid.Coord
}

// map of position visited and all the different ways
// we were facing when we visited it.
type Visited map[grid.Coord][]grid.Coord

func Solve(file string, partOne bool) int {
	var start grid.Coord

	maze := grid.Parse(utils.MustReadInput(file), func(value rune, x, y int) bool {
		if value == '^' {
			start = grid.At(x, y)
			return false
		}
		return value == '#'
	})

	path, loop := simulate(&maze, start, grid.North)
	assert.False(loop, "we shouldn't have any loops in base path")

	if partOne {
		return len(path)
	}

	loops := utils.NewSet[grid.Coord]()

	// iterate over the path we walked to complete the maze and place walls
	// ahead the direction we were moving to test to see if this would have
	// resulted in a loop
	for node, facings := range path {
		for _, facing := range facings {
			wallPosition := node.Offset(facing)

			if !maze.Contains(wallPosition) {
				continue
			}
			if loops.Contains(wallPosition) {
				continue
			}
			if isWall, _ := maze.GetContents(wallPosition); isWall {
				continue
			}
			if wallPosition == start {
				continue
			}

			maze.Replace(wallPosition, true)
			if _, looped := simulate(&maze, start, grid.North); looped {
				loops.Add(wallPosition)
			}
			maze.Replace(wallPosition, false)
		}
	}

	return len(loops)
}

func simulate(maze *grid.Grid[bool], position grid.Coord, facing grid.Coord) (Visited, bool) {
	visited := make(Visited)
	visited[position] = []grid.Coord{facing}

	for {
		if !maze.Contains(position) {
			delete(visited, position)
			return visited, false
		}

		next := position.Offset(facing)
		if isWall, _ := maze.GetContents(next); isWall {
			facing = facing.TurnRight()
			visited[position] = append(visited[position], facing)
			continue
		}

		position = next

		facings := visited[position]
		if slices.Contains(facings, facing) {
			return visited, true
		}

		visited[position] = append(visited[position], facing)
	}
}
