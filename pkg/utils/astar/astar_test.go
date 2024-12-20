package astar

import (
	"adventofcode2024/pkg/utils/grid"
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueueFunctions(t *testing.T) {
	pq := make(PriorityQueue[grid.Coordinate], 0)
	heap.Init(&pq)

	heap.Push(&pq, NewNode(grid.At(2, 0), 2, nil))
	heap.Push(&pq, NewNode(grid.At(3, 0), 3, nil))
	heap.Push(&pq, NewNode(grid.At(1, 0), 1, nil))
	heap.Push(&pq, NewNode(grid.At(0, 0), 0, nil))

	assert.Equal(t, 4, pq.Len())
	assert.Equal(t, grid.At(0, 0), heap.Pop(&pq).(*Node[grid.Coordinate]).Contents)
	assert.Equal(t, grid.At(1, 0), heap.Pop(&pq).(*Node[grid.Coordinate]).Contents)
	assert.Equal(t, grid.At(2, 0), heap.Pop(&pq).(*Node[grid.Coordinate]).Contents)
	assert.Equal(t, grid.At(3, 0), heap.Pop(&pq).(*Node[grid.Coordinate]).Contents)
}

func TestCanSolvePath(t *testing.T) {
	start := grid.At(0, 0)
	end := grid.At(1, 1)
	neighbors := func(node *Node[grid.Coordinate]) []grid.Coordinate {
		source := node.Contents
		return []grid.Coordinate{source.North(), source.South(), source.East(), source.West()}
	}

	heuristic := func(node grid.Coordinate, from *Node[grid.Coordinate]) float64 {
		return float64(node.Distance(end))
	}

	path := Find(start, end, neighbors, heuristic)
	assert.Equal(t, 3, len(path))
	assert.Equal(t, start, path[0])
	// assert.Equal(t, grid.At(1, 0), path[1])
	assert.Equal(t, end, path[2])
}
