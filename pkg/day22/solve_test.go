package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 37327623, PartOne("input.example.one.txt"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 13004408787, PartOne("input.txt"))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 23, PartTwo("input.example.two.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 1455, PartTwo("input.txt"))
}
