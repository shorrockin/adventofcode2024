package day10

import (
	"adventofcode2024/pkg/grid"
	"adventofcode2024/pkg/utils"
)

func Solve(path string, partOne bool) int {
	trailheads := make([]grid.Coordinate, 0)
	guide := grid.Parse(utils.MustReadInput(path), func(value rune, x, y int) int {
		if value == '0' {
			trailheads = append(trailheads, grid.At(x, y))
		}
		return utils.MustAtoi(string(value))
	})

	return bfs(guide, trailheads, partOne)
}

func bfs(guide grid.Grid[int], trailheads []grid.Coordinate, partOne bool) int {
	found := 0

	for _, trailhead := range trailheads {
		queue := utils.NewQueue[[]grid.Coordinate]()
		queue.Enqueue([]grid.Coordinate{trailhead})
		visited := utils.NewSet[string]()

		for !queue.IsEmpty() {
			current := queue.MustDequeue()
			tail := current[len(current)-1]

			if guide.MustGetContents(tail) == 9 {
				found++
				continue
			}

			for _, neighbor := range neighbors(&guide, tail) {
				var path []grid.Coordinate = make([]grid.Coordinate, len(current)+1)
				copy(path, current)
				path[len(current)] = neighbor
				hash := generateHash(&path, partOne)

				if !visited.Contains(hash) {
					visited.Add(hash)
					queue.Enqueue(path)
				}
			}
		}
	}

	return found
}

func neighbors(guide *grid.Grid[int], from grid.Coordinate) []grid.Coordinate {
	base := guide.MustGetContents(from)
	return utils.Filter(from.Cardinals(), func(neighbor grid.Coordinate) bool {
		if value, exists := guide.GetContents(neighbor); exists {
			return (value - base) == 1
		}
		return false
	})
}

func generateHash(path *[]grid.Coordinate, partOne bool) string {
	if partOne {
		return (*path)[len(*path)-1].String()
	}

	hash := ""
	for _, coordinate := range *path {
		hash += coordinate.String()
	}
	return hash
}
