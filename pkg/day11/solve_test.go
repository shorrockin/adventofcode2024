package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 55312, Solve("125 17", 25))
}

func TestSplitting(t *testing.T) {
	left, right, ok := split(123456)
	assert.Equal(t, 123, left)
	assert.Equal(t, 456, right)
	assert.True(t, ok)

	left, right, ok = split(3000)
	assert.Equal(t, 30, left)
	assert.Equal(t, 0, right)
	assert.True(t, ok)

	left, right, ok = split(123)
	assert.Equal(t, 0, left)
	assert.Equal(t, 0, right)
	assert.False(t, ok)
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 207683, Solve("3935565 31753 437818 7697 5 38 0 123", 25))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 244782991106220, Solve("3935565 31753 437818 7697 5 38 0 123", 75))
}
