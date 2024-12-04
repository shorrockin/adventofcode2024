package day04

import (
	"adventofcode2024/pkg/grid"
	"adventofcode2024/pkg/utils"
	"slices"
)

func PartOne(path string) int {
	wordSearch := parse(path)
	found := 0

	for _, node := range wordSearch {
		for _, direction := range grid.Directions {
			x := node.Coordinate
			m := x.Offset(direction)
			a := m.Offset(direction)
			s := a.Offset(direction)

			if letters, ok := get(wordSearch, x, m, a, s); ok {
				if string(letters) == "XMAS" {
					found++
				}
			}
		}
	}
	return found
}

func PartTwo(path string) int {
	wordSearch := parse(path)
	found := 0

	for _, node := range wordSearch {
		center := node.Coordinate
		nw := center.Offset(grid.NorthWest)
		ne := center.Offset(grid.NorthEast)
		sw := center.Offset(grid.SouthWest)
		se := center.Offset(grid.SouthEast)

		if letters, ok := get(wordSearch, center, nw, ne, sw, se); ok {
			if letters[1] == letters[4] {
				// opposing corners shouldn't match
				continue
			}
			if letters[0] != 'A' {
				// 'A' should be in the center
				continue
			}
			slices.Sort(letters)
			if string(letters) != "AMMSS" {
				// remaining letters can only be M or S, two of each
				continue
			}

			found++
		}
	}
	return found
}

func parse(path string) grid.Grid[rune] {
	return grid.Parse(utils.MustReadInput(path), func(value rune, x int, y int) rune {
		return value
	})
}

func get(search grid.Grid[rune], coordinates ...grid.Coordinate) ([]rune, bool) {
	var contents = make([]rune, len(coordinates))
	for idx, coordinate := range coordinates {
		if node, ok := search.Get(coordinate); ok {
			contents[idx] = node.Contents
		} else {
			return nil, false
		}
	}
	return contents, true
}
