package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 2, Solve("input.example.txt", false))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 4, Solve("input.example.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 369, Solve("input.txt", false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 428, Solve("input.txt", true))
}

// func TestGaaahConfusion(t *testing.T) {
// 	ints := []int{0, 1, 2, 3}
// 	modified := append(ints[:1], ints[2:]...)
//
// 	assert.Equal(t, []int{0, 2, 3}, modified)
//
//  this fails...
// 	assert.Equal(t, []int{0, 1, 2, 3}, ints)
// }
