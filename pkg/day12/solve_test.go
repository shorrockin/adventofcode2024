package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExamples(t *testing.T) {
	assert.Equal(t, 140, Solve("input.example-one.txt", true))
	assert.Equal(t, 772, Solve("input.example-two.txt", true))
	assert.Equal(t, 1930, Solve("input.example-three.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 1431440, Solve("input.txt", true))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 80, Solve("input.example-one.txt", false))
	assert.Equal(t, 436, Solve("input.example-two.txt", false))
	assert.Equal(t, 236, Solve("input.example-four.txt", false))
	assert.Equal(t, 368, Solve("input.example-five.txt", false))
	assert.Equal(t, 1206, Solve("input.example-three.txt", false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 869070, Solve("input.txt", false))
}
