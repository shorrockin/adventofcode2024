package day10

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/collections"
	"adventofcode2024/pkg/utils/grid"
)

func Solve(path string, partOne bool) int {
	trailheads := make([]grid.Coord, 0)
	guide := grid.Parse(utils.MustReadInput(path), func(value rune, x, y int) int {
		if value == '0' {
			trailheads = append(trailheads, grid.At(x, y))
		}
		return utils.MustAtoi(string(value))
	})

	return bfs(guide, trailheads, partOne)
}

func bfs(guide grid.Grid[int], trailheads []grid.Coord, partOne bool) int {
	found := 0

	for _, trailhead := range trailheads {
		queue := collections.NewQueue[[]grid.Coord]()
		queue.Push([]grid.Coord{trailhead})
		visited := collections.NewSet[string]()

		for !queue.IsEmpty() {
			current := queue.MustPop()
			tail := current[len(current)-1]

			if guide.MustGetContents(tail) == 9 {
				found++
				continue
			}

			for _, neighbor := range neighbors(&guide, tail) {
				var path []grid.Coord = make([]grid.Coord, len(current)+1)
				copy(path, current)
				path[len(current)] = neighbor
				hash := generateHash(&path, partOne)

				if !visited.Contains(hash) {
					visited.Add(hash)
					queue.Push(path)
				}
			}
		}
	}

	return found
}

func neighbors(guide *grid.Grid[int], from grid.Coord) []grid.Coord {
	base := guide.MustGetContents(from)
	return utils.Filter(from.Cardinals(), func(neighbor grid.Coord) bool {
		if value, exists := guide.GetContents(neighbor); exists {
			return (value - base) == 1
		}
		return false
	})
}

func generateHash(path *[]grid.Coord, partOne bool) string {
	if partOne {
		return (*path)[len(*path)-1].String()
	}

	hash := ""
	for _, coordinate := range *path {
		hash += coordinate.String()
	}
	return hash
}
