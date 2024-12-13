package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 480, Solve("input.example.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 32041, Solve("input.txt", true))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 875318608908, Solve("input.example.txt", false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 95843948914827, Solve("input.txt", false))
}
