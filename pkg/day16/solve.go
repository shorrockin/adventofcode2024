package day16

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/collections"
	"adventofcode2024/pkg/utils/grid"
)

type Movement struct {
	position grid.Coord
	facing   grid.Coord
}

type Score struct {
	value int
	path  collections.Set[grid.Coord]
}

func Solve(path string, partOne bool) int {
	start := grid.At(-1, -1)
	end := grid.At(-1, -1)
	maze := grid.Parse(utils.MustReadInput(path), func(value rune, x, y int) bool {
		if value == '#' {
			return false
		} else if value == 'S' {
			start = grid.At(x, y)
		} else if value == 'E' {
			end = grid.At(x, y)
		}

		return true
	})

	queue := collections.NewQueue[Movement]()
	queue.Push(Movement{start, grid.East})

	scores := make(map[Movement]Score)
	scores[Movement{start, grid.East}] = Score{0, collections.NewSetFrom(start)}

	for !queue.IsEmpty() {
		origin := queue.MustPop()
		movements := []Movement{
			{origin.position.Offset(origin.facing), origin.facing},
			{origin.position, origin.facing.TurnLeft()},
			{origin.position, origin.facing.TurnRight()},
		}

		for idx, movement := range movements {
			walkable, exists := maze.GetContents(movement.position)
			if !exists || !walkable {
				continue
			}

			score := scores[origin].value + 1
			if idx != 0 {
				score = scores[origin].value + 1000
			}

			if best, exists := scores[movement]; !exists || score <= best.value {
				if score == best.value {
					for cord := range scores[origin].path {
						best.path.Add(cord)
					}
					best.path.Add(movement.position)
					continue
				}
				scores[movement] = scores[origin].step(score, movement.position)
				queue.Push(movement)
			}
		}
	}

	score := findScoreAt(scores, end)
	if partOne {
		return score.value
	}
	return len(score.path)
}

func (s Score) step(value int, coord grid.Coord) Score {
	path := s.path.Copy()
	path.Add(coord)
	return Score{value, path}
}

func findScoreAt(scores map[Movement]Score, position grid.Coord) Score {
	best := Score{}
	for _, movement := range []Movement{
		{position, grid.East},
		{position, grid.West},
		{position, grid.North},
		{position, grid.South},
	} {
		if score, exists := scores[movement]; exists && (best.value == 0 || score.value < best.value) {
			best = score
		}
	}
	return best
}
