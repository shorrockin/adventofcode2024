package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 126384, Solve("input.example.txt", 3))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 154208, Solve("input.txt", 3))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 188000493837892, Solve("input.txt", 26))
}
