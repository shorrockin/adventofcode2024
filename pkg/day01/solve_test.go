package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 11, Solve("input.example.txt", PartOne))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 31, Solve("input.example.txt", PartTwo))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 1646452, Solve("input.txt", PartOne))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 23609874, Solve("input.txt", PartTwo))
}
