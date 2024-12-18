package day18

import (
	"adventofcode2024/pkg/grid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 22, PartOne("input.example.txt", 6, 6, 12))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 288, PartOne("input.txt", 70, 70, 1024))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, grid.At(6, 1), PartTwo("input.example.txt", 6, 6, 12))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, grid.At(52, 5), PartTwo("input.txt", 70, 70, 1024))
}
