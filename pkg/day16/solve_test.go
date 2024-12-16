package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExampleOne(t *testing.T) {
	assert.Equal(t, 7036, Solve("input.example.one.txt", true))
}

func TestPartOneExampleTwo(t *testing.T) {
	assert.Equal(t, 11048, Solve("input.example.two.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 106512, Solve("input.txt", true))
}

func TestPartTwoExampleOne(t *testing.T) {
	assert.Equal(t, 45, Solve("input.example.one.txt", false))
}

func TestParTwoExampleTwo(t *testing.T) {
	assert.Equal(t, 64, Solve("input.example.two.txt", false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 563, Solve("input.txt", false))
}
