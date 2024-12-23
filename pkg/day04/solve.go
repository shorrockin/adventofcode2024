package day04

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/grid"
	"adventofcode2024/pkg/utils/slices"
)

type Pattern struct {
	offsets []grid.Coord
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
			coordinates := slices.Map(pattern.offsets, func(c grid.Coord) grid.Coord {
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
		return slices.Map(grid.Directions, func(direction grid.Coord) Pattern {
			x := grid.At(0, 0)
			m := x.Offset(direction)
			a := m.Offset(direction)
			s := a.Offset(direction)
			return Pattern{[]grid.Coord{x, m, a, s}, "XMAS"}
		})
	}

	return slices.Map([]string{"AMMSS", "ASSMM", "AMSMS", "ASMSM"}, func(value string) Pattern {
		return Pattern{[]grid.Coord{grid.At(0, 0), grid.NorthWest, grid.NorthEast, grid.SouthWest, grid.SouthEast}, value}
	})
}

func extract(wordSearch grid.Grid[rune], coordinates ...grid.Coord) ([]rune, bool) {
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
