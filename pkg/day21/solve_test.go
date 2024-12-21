package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	// assert.Equal(t, 126384, Solve("input.example.txt", 3))
}

func TestPartOneActual(t *testing.T) {
	// assert.Equal(t, "X", BestPath(&numerics, 'A', '0', 2))
	assert.Equal(t, -154208, Solve("input.txt", 22))
}

// func TestPartTwoExample(t *testing.T) {
// 	assert.Equal(t, -1, Solve("input.example.txt", false))
// }

func TestPartTwoActual(t *testing.T) {
	// assert.Equal(t, -1, Solve("input.txt", 25))
}
