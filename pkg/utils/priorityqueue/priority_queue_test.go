package priorityqueue

import (
	"adventofcode2024/pkg/utils/grid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueueFunctions(t *testing.T) {
	pq := NewPriorityQueue[grid.Coord]()

	pq.Push(grid.At(2, 0), 2, nil)
	pq.Push(grid.At(3, 0), 3, nil)
	pq.Push(grid.At(1, 0), 1, nil)
	pq.Push(grid.At(0, 0), 0, nil)

	assert.Equal(t, 4, pq.Len())
	assert.Equal(t, grid.At(0, 0), pq.Pop())
	assert.Equal(t, grid.At(1, 0), pq.Pop())
	assert.Equal(t, grid.At(2, 0), pq.Pop())
	assert.Equal(t, grid.At(3, 0), pq.Pop())
}
