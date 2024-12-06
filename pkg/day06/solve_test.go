package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 41, Solve("input.example.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 4776, Solve("input.txt", true))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 6, Solve("input.example.txt", false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 1586, Solve("input.txt", false))
}
