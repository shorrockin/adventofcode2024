package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 14, Solve("input.example.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 280, Solve("input.txt", true))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 34, Solve("input.example.txt", false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 958, Solve("input.txt", false))
}
