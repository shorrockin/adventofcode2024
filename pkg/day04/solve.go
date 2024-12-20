package day04

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/grid"
)

type Pattern struct {
	offsets []grid.Coordinate
	value   string
}

func Solve(path string, partTwo bool) int {
	found := 0
	patterns := generatePatterns(partTwo)
	wordSearch := grid.Parse(utils.MustReadInput(path), func(value rune, x int, y int) rune {
		return value
	})

	for _, node := range wordSearch {
		for _, pattern := range patterns {
			coordinates := utils.Map(pattern.offsets, func(c grid.Coordinate) grid.Coordinate {
				return node.Coordinate.Offset(c)
			})

			if vlaue, ok := extract(wordSearch, coordinates...); ok {
				if string(vlaue) == pattern.value {
					found++
				}
			}
		}
	}
	return found
}

func generatePatterns(partTwo bool) []Pattern {
	if !partTwo {
		return utils.Map(grid.Directions, func(direction grid.Coordinate) Pattern {
			x := grid.At(0, 0)
			m := x.Offset(direction)
			a := m.Offset(direction)
			s := a.Offset(direction)
			return Pattern{[]grid.Coordinate{x, m, a, s}, "XMAS"}
		})
	}

	return utils.Map([]string{"AMMSS", "ASSMM", "AMSMS", "ASMSM"}, func(value string) Pattern {
		return Pattern{[]grid.Coordinate{grid.At(0, 0), grid.NorthWest, grid.NorthEast, grid.SouthWest, grid.SouthEast}, value}
	})
}

func extract(wordSearch grid.Grid[rune], coordinates ...grid.Coordinate) ([]rune, bool) {
	var contents = make([]rune, len(coordinates))

	for idx, coordinate := range coordinates {
		if node, ok := wordSearch.Get(coordinate); ok {
			contents[idx] = node.Contents
		} else {
			return nil, false
		}
	}

	return contents, true
}
