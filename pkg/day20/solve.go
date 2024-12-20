package day20

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/graph"
	"adventofcode2024/pkg/utils/graph/dijkstra"
	"adventofcode2024/pkg/utils/grid"
)

func Solve(path string, savings, cheatRange int) int {
	maze, start := parse(path)
	cheats := 0
	distances := dijkstra.Dijkstra(maze, start)

	for from, fromDistance := range distances {
		for _, to := range from.CoordsInRange(cheatRange) {
			if toDistance, exists := distances[to]; exists && fromDistance < toDistance {
				if (int(toDistance) - int(fromDistance) - from.Distance(to)) >= savings {
					cheats++
				}
			}
		}
	}

	return cheats
}

func parse(path string) (*graph.Graph[grid.Coord], grid.Coord) {
	var start grid.Coord

	maze := grid.Parse(utils.MustReadInput(path), func(value rune, x, y int) bool {
		if value == 'S' {
			start = grid.At(x, y)
		}
		return value == '#'
	})

	return maze.ToGraph(func(node grid.Node[bool]) bool {
		return !node.Contents
	}), start
}
