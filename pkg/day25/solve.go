package day25

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/slices"
)

type Schematic struct {
	idx     int
	heights [5]int
	key     bool
}

func Solve(path string, partOne bool) int {
	lines := utils.MustReadInput(path)
	schematics := make([]Schematic, 0)

	for idx := 0; idx < len(lines); idx++ {
		if lines[idx] == "....." || lines[idx] == "#####" {
			var heights [5]int
			for col := 0; col < 5; col++ {
				for row := 0; row < 5; row++ {
					if lines[idx+row+1][col] == '#' {
						heights[col]++
					}
				}
			}
			schematics = append(schematics, Schematic{idx, heights, lines[idx] == "....."})
			idx += 7
		}
	}

	count := 0
	for _, combo := range slices.Combinations(schematics, 2) {
		if combo[0].key == combo[1].key {
			continue
		}

		for col := 0; col < 5; col++ {
			if combo[0].heights[col]+combo[1].heights[col] > 5 {
				break
			} else if col == 4 {
				count++
			}
		}
	}

	return count
}
