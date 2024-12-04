package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 18, PartOne("input.example.txt"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 2427, PartOne("input.txt"))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 9, PartTwo("input.example.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 1900, PartTwo("input.txt"))
}
