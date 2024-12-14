package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 12, Solve("input.example.txt", 11, 7, true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 225552000, Solve("input.txt", 101, 103, true))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 7371, Solve("input.txt", 101, 103, false))
}
