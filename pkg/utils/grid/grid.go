package grid

import (
	"adventofcode2024/pkg/assert"
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/graph"
	"math"
	"strings"
)

type Node[T any] struct {
	Coordinate Coordinate
	Contents   T
}

type Grid[T any] map[Coordinate]Node[T]

func NewGrid[T any]() Grid[T] {
	return make(Grid[T])
}

func Parse[T any](lines []string, init func(value rune, x int, y int) T) Grid[T] {
	grid := NewGrid[T]()
	for y, line := range lines {
		for x, char := range line {
			grid.InsertAt(x, y, init(char, x, y))
		}
	}
	return grid
}

func (g Grid[T]) At(x int, y int) (Node[T], bool) {
	return g.Get(At(x, y))
}

func (g Grid[T]) MustAt(x int, y int) Node[T] {
	return g.MustGet(At(x, y))
}

func (g Grid[T]) Get(coordinate Coordinate) (Node[T], bool) {
	value, ok := g[coordinate]
	return value, ok
}

func (g Grid[T]) GetContents(coordinate Coordinate) (T, bool) {
	value, ok := g[coordinate]
	return value.Contents, ok
}

func (g Grid[T]) MustGetContents(coordinate Coordinate) T {
	g.AssertPopulated(coordinate)
	return g[coordinate].Contents
}

func (g Grid[T]) Contains(coordinate Coordinate) bool {
	_, ok := g[coordinate]
	return ok
}

func (g Grid[T]) MustGet(coordinate Coordinate) Node[T] {
	g.AssertPopulated(coordinate)
	return g[coordinate]
}

func (g Grid[T]) GetAll(coordinates []Coordinate) []Node[T] {
	values := utils.MapConditional(coordinates, func(coordinate Coordinate) (Node[T], bool) {
		value, ok := g[coordinate]
		return value, ok
	})

	return values
}

func (g Grid[T]) InsertAt(x, y int, value T) {
	g.Insert(At(x, y), value)
}

func (g Grid[T]) Insert(coordinate Coordinate, value T) {
	g.AssertEmpty(coordinate)
	g[coordinate] = Node[T]{
		Coordinate: coordinate,
		Contents:   value,
	}
}

func (g Grid[T]) ReplaceAt(x, y int, value T) {
	g.Replace(At(x, y), value)
}

func (g Grid[T]) Replace(coordinate Coordinate, value T) {
	g[coordinate] = Node[T]{
		Coordinate: coordinate,
		Contents:   value,
	}
}

func (g Grid[T]) String(empty rune, mapper func(contents T) rune) string {
	var builder strings.Builder
	minWidth, maxWidth := g.Width()
	minHeight, maxHeight := g.Height()

	for y := minHeight; y <= maxHeight; y++ {
		if y != minHeight {
			builder.WriteString("\n")
		}
		for x := minWidth; x <= maxWidth; x++ {
			node, ok := g.At(x, y)
			if !ok {
				builder.WriteRune(empty)
				continue
			}
			builder.WriteRune(mapper(node.Contents))
		}
	}
	return builder.String()
}

func (g Grid[T]) Width() (int, int) {
	max := math.MinInt
	min := math.MaxInt
	for coordinate := range g {
		max = utils.MaxValue(coordinate.X, max)
		min = utils.MinValue(coordinate.X, min)
	}
	return min, max
}

func (g Grid[T]) Height() (int, int) {
	max := math.MinInt
	min := math.MaxInt
	for coordinate := range g {
		max = utils.MaxValue(coordinate.Y, max)
		min = utils.MinValue(coordinate.Y, min)
	}
	return min, max
}

func (g Grid[T]) UpdateAt(coordinate Coordinate, value T) {
	g.AssertPopulated(coordinate)
	g[coordinate] = Node[T]{coordinate, value}
}

func (g Grid[T]) AssertPopulated(coordinate Coordinate) {
	_, ok := g[coordinate]
	assert.True(ok, "expected value to be populated at coordinate", coordinate)
}

func (g Grid[T]) AssertEmpty(coordinate Coordinate) {
	_, ok := g[coordinate]
	assert.False(ok, "expected value to not be populated at coordinate", "coordinate", coordinate)
}

func (g Grid[T]) ToGraph(validator func(node Node[T]) bool) *graph.Graph[Coordinate] {
	grph := graph.NewGraph[Coordinate]()
	for coordinate, node := range g {
		if !validator(node) {
			continue
		}

		for _, neighbor := range coordinate.Cardinals() {
			if neighborNode, ok := g.Get(neighbor); ok && validator(neighborNode) {
				grph.AddBidirectionalEdge(coordinate, neighbor, 1)
			}
		}
	}
	return grph
}
