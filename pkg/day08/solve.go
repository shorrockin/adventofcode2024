package day08

import (
	"adventofcode2024/pkg/assert"
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/grid"
)

func Solve(path string, partOne bool) int {
	lines := utils.MustReadInput(path)
	height := len(lines)
	width := len(lines[0])
	attenas := make(map[rune][]grid.Coord)
	antinodes := utils.NewSet[grid.Coord]()

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				attenas[char] = append(attenas[char], grid.At(x, y))
			}
		}
	}

	for _, coordinates := range attenas {
		combinations := utils.Combinations(coordinates, 2)

		for _, combination := range combinations {
			assert.Equal(2, len(combination), "there should only be 2 elements", combination)

			for idx, self := range combination {
				other := combination[(idx+1)%2]
				growthX := self.X - other.X
				growthY := self.Y - other.Y

				if !partOne {
					antinodes.Add(self)
				}

				for {
					antinode := grid.At(self.X+growthX, self.Y+growthY)
					if !inBounds(antinode, width, height) {
						break
					}

					antinodes.Add(antinode)

					if partOne {
						break
					}
					growthX += (self.X - other.X)
					growthY += (self.Y - other.Y)
				}
			}
		}
	}

	return len(antinodes)
}

func inBounds(position grid.Coord, width, height int) bool {
	return position.X >= 0 &&
		position.Y >= 0 &&
		position.X < width &&
		position.Y < height
}
