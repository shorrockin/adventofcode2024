package day16

import (
	"adventofcode2024/pkg/grid"
	"adventofcode2024/pkg/utils"
)

type Movement struct {
	position grid.Coordinate
	facing   grid.Coordinate
}

type Score struct {
	value int
	path  utils.Set[grid.Coordinate]
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

	queue := utils.NewQueue[Movement]()
	queue.Enqueue(Movement{start, grid.East})

	scores := make(map[Movement]Score)
	scores[Movement{start, grid.East}] = Score{0, utils.NewSetFrom(start)}

	for !queue.IsEmpty() {
		source := queue.MustDequeue()
		movements := []Movement{
			{source.position.Offset(source.facing), source.facing},
			{source.position, source.facing.TurnLeft()},
			{source.position, source.facing.TurnRight()},
		}

		for idx, movement := range movements {
			walkable, exists := maze.GetContents(movement.position)
			if !exists || !walkable {
				continue
			}

			score := scores[source].spawn(1, movement.position)
			if idx != 0 {
				score.value = scores[source].value + 1000
			}

			if best, exists := scores[movement]; !exists || score.value <= best.value {
				if score.value == best.value {
					scores[movement].merge(score)
					continue
				}
				scores[movement] = score
				queue.Enqueue(movement)
			}
		}
	}

	score := findScoreAt(&scores, end)
	if partOne {
		return score.value
	}
	return len(score.path)
}

func (s Score) merge(other Score) {
	for _, cord := range other.path.Values() {
		s.path.Add(cord)
	}
}

func (s Score) spawn(value int, coord grid.Coordinate) Score {
	path := s.path.Clone()
	path.Add(coord)
	return Score{s.value + value, path}
}

func findScoreAt(scores *map[Movement]Score, position grid.Coordinate) Score {
	best := Score{}
	for _, movement := range []Movement{
		{position, grid.East},
		{position, grid.West},
		{position, grid.North},
		{position, grid.South},
	} {
		if score, exists := (*scores)[movement]; exists && (best.value == 0 || score.value < best.value) {
			best = score
		}
	}
	return best
}
