package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 7, PartOne("input.example.txt"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 1323, PartOne("input.txt"))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, "co,de,ka,ta", PartTwo("input.example.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, "er,fh,fi,ir,kk,lo,lp,qi,ti,vb,xf,ys,yu", PartTwo("input.txt"))
}
