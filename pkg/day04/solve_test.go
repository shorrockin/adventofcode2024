package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 18, Solve("input.example.txt", false))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 2427, Solve("input.txt", false))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 9, Solve("input.example.txt", true))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 1900, Solve("input.txt", true))
}
