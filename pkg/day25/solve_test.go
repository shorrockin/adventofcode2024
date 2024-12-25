package day25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 3, Solve("input.example.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 3466, Solve("input.txt", true))
}

// func TestPartTwoExample(t *testing.T) {
// 	assert.Equal(t, -1, Solve("input.example.txt", false))
// }

// func TestPartTwoActual(t *testing.T) {
// 	assert.Equal(t, -1, Solve("input.txt", false))
// }
