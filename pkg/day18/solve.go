package day18

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/assert"
	"adventofcode2024/pkg/utils/astar"
	"adventofcode2024/pkg/utils/grid"
	"strings"
)

func PartOne(path string, width, height, time int) int {
	return ShortestPath(parse(path), width, height, time)
}

func PartTwo(path string, width, height, startTime int) grid.Coord {
	bytes := parse(path)
	time := startTime

	for ShortestPath(bytes, width, height, time) > 0 {
		time++
	}

	for coord, at := range bytes {
		if at == (time - 1) {
			return coord
		}
	}

	panic(assert.Fail("could not find blocking path"))
}

func ShortestPath(bytes map[grid.Coord]int, width, height, time int) int {
	end := grid.At(width, height)

	neighbors := func(from *astar.Node[grid.Coord]) []grid.Coord {
		return utils.Filter(from.Contents.Cardinals(), func(pos grid.Coord) bool {
			if byteTime, ok := bytes[pos]; ok && byteTime < time {
				return false
			}

			if pos.X < 0 || pos.X > width || pos.Y < 0 || pos.Y > height {
				return false
			}

			return true
		})
	}

	heuristic := func(node grid.Coord, from *astar.Node[grid.Coord]) float64 {
		return float64(from.PathDepth) + float64(node.Distance(end))
	}

	solution := astar.Find(grid.At(0, 0), end, neighbors, heuristic)
	return len(solution) - 1
}

func parse(path string) map[grid.Coord]int {
	bytes := utils.Map(utils.MustReadInput(path), func(line string) grid.Coord {
		parts := strings.Split(line, ",")
		assert.Equal(2, len(parts), "expected two parts to each coordinate")
		return grid.Coord{X: utils.MustAtoi(parts[0]), Y: utils.MustAtoi(parts[1])}
	})

	time := make(map[grid.Coord]int)
	for idx, coord := range bytes {
		time[coord] = idx
	}
	return time
}
