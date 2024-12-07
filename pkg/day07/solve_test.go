package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 3749, Solve("input.example.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 20281182715321, Solve("input.txt", true))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 11387, Solve("input.example.txt", false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 159490400628354, Solve("input.txt", false))
}
