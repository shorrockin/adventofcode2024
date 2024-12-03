package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 161, Solve("input.example.one.txt", PartOne))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 48, Solve("input.example.two.txt", PartTwo))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 161289189, Solve("input.txt", PartOne))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 83595109, Solve("input.txt", PartTwo))
}
