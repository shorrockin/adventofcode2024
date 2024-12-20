package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 44, Solve("input.example.txt", 2, 2))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 1485, Solve("input.txt", 100, 2))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 29, Solve("input.example.txt", 72, 20))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 1027501, Solve("input.txt", 100, 20))
}
