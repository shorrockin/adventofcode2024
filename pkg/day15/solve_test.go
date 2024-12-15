package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExampleBig(t *testing.T) {
	assert.Equal(t, 10092, Solve("input.example.big.txt", true))
}

func TestPartOneExampleSmall(t *testing.T) {
	assert.Equal(t, 2028, Solve("input.example.small.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 1446158, Solve("input.txt", true))
}

func TestPartTwoExampleSmall(t *testing.T) {
	assert.Equal(t, 618, Solve("input.example.small.two.txt", false))
}

func TestPartTwoExampleBig(t *testing.T) {
	assert.Equal(t, 9021, Solve("input.example.big.txt", false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 1446175, Solve("input.txt", false))
}
