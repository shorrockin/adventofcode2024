package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 143, Solve("input.example.txt", PartOne))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 5064, Solve("input.txt", PartOne))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 123, Solve("input.example.txt", PartTwo))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 5152, Solve("input.txt", PartTwo))
}
