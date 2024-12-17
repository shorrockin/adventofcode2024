package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}, PartOne("input.example.one.txt"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, []int{1, 5, 0, 3, 7, 3, 0, 3, 1}, PartOne("input.txt"))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 117440, PartTwo("input.example.two.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, -105981155568026, PartTwo("input.txt"))
}
