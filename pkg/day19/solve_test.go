package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 6, Solve("input.example.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 242, Solve("input.txt", true))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 16, Solve("input.example.txt", false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 595975512785325, Solve("input.txt", false))
}
