package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 11, Solve("input.example.txt", Part1))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 31, Solve("input.example.txt", Part2))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 1646452, Solve("input.txt", Part1))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 23609874, Solve("input.txt", Part2))
}
