package astar

import (
	"adventofcode2024/pkg/utils/grid"
	"adventofcode2024/pkg/utils/priorityqueue"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSolvePath(t *testing.T) {
	start := grid.At(0, 0)
	end := grid.At(1, 1)
	neighbors := func(node *priorityqueue.Node[grid.Coord]) []grid.Coord {
		source := node.Contents
		return []grid.Coord{source.North(), source.South(), source.East(), source.West()}
	}

	heuristic := func(node grid.Coord, from *priorityqueue.Node[grid.Coord]) float64 {
		return float64(node.Distance(end))
	}

	path := AStar(start, end, neighbors, heuristic)
	assert.Equal(t, 3, len(path))
	assert.Equal(t, start, path[0])
	// assert.Equal(t, grid.At(1, 0), path[1])
	assert.Equal(t, end, path[2])
}
