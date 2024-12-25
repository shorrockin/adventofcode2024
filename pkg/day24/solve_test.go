package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestPartOneExample(t *testing.T) {
// 	assert.Equal(t, 2024, PartOne("input.example.txt"))
// }
//
// func TestPartOneActual(t *testing.T) {
// 	assert.Equal(t, 51837135476040, PartOne("input.txt"))
// }

// func TestPartTwoExample(t *testing.T) {
// 	assert.Equal(t, -1, PartTwo("input.example.two.txt"))
// }

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, "hjf,kdh,kpp,sgj,vss,z14,z31,z35", Original("input.txt"))
	// assert.Equal(t, "hjf,kdh,kpp,sgj,vss,z14,z31,z35", PartTwo("input.txt"))
}
