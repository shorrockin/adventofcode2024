package day15

import (
	"adventofcode2024/pkg/assert"
	"adventofcode2024/pkg/grid"
	"adventofcode2024/pkg/utils"
	"fmt"
	"strings"
)

type Item int

const (
	Robot Item = iota
	Wall
	Box
	BoxLeft
	BoxRight
	Open
)

type Warehouse struct {
	layout    grid.Grid[Item]
	movements utils.Queue[grid.Coordinate]
	robot     grid.Coordinate
}

func Solve(path string, partOne bool) int {
	warehouse := parse(path, partOne)

	for !warehouse.movements.IsEmpty() {
		direction := warehouse.movements.MustDequeue()
		warehouse.robot = move(warehouse, warehouse.robot, direction)
	}

	return score(warehouse)
}

func move(warehouse *Warehouse, from grid.Coordinate, direction grid.Coordinate) grid.Coordinate {
	dest := from.Offset(direction)
	current := warehouse.layout.MustGetContents(from)

	execute := func(from grid.Coordinate, dest grid.Coordinate) grid.Coordinate {
		warehouse.layout.UpdateAt(dest, current)
		warehouse.layout.UpdateAt(from, Open)
		return dest
	}

	if destItem, ok := warehouse.layout.GetContents(dest); ok {
		if destItem == Open {
			// open space - move into it
			return execute(from, dest)
		} else if destItem == Box {
			// part 1 box - recurse and move in if the box was able to move
			if dest != move(warehouse, dest, direction) {
				return execute(from, dest)
			}
		} else if (destItem == BoxLeft || destItem == BoxRight) && (direction == grid.East || direction == grid.West) {
			// part 2 box, moving left / right - same as above, recurse and move if we can
			if dest != move(warehouse, dest, direction) {
				return execute(from, dest)
			}
		} else if destItem == BoxLeft || destItem == BoxRight {
			// part 2 box moving up/down: need to move together only if open space exists
			// above, need to check for cascading open space on paired boxes
			otherHalf := dest.Offset(pairedBox(destItem))
			if canMoveAll(warehouse, direction, dest, otherHalf) {
				move(warehouse, dest, direction)
				move(warehouse, otherHalf, direction)
				return execute(from, dest)
			}

		}
	}

	return from
}

func canMoveAll(warehouse *Warehouse, direction grid.Coordinate, dests ...grid.Coordinate) bool {
	for _, dest := range dests {
		target := dest.Offset(direction)
		item := warehouse.layout.MustGetContents(target)

		if item == Wall {
			return false
		} else if item == BoxLeft || item == BoxRight {
			if !canMoveAll(warehouse, direction, target, target.Offset(pairedBox(item))) {
				return false
			}
		}
	}
	return true
}

func pairedBox(item Item) grid.Coordinate {
	switch item {
	case BoxLeft:
		return grid.East
	case BoxRight:
		return grid.West
	default:
		panic(assert.Fail("unable to find box buddy for invalid item", "item", item))
	}
}

func score(warehouse *Warehouse) int {
	sum := 0
	for coordinate, node := range warehouse.layout {
		if node.Contents == Box || node.Contents == BoxLeft {
			sum += (100*coordinate.Y + coordinate.X)
		}
	}
	return sum
}

func debug(warehouse *Warehouse) string {
	return fmt.Sprintf("%v", warehouse.layout.String('x', func(contents Item) rune {
		switch contents {
		case Robot:
			return '@'
		case Wall:
			return '#'
		case Box:
			return 'O'
		case Open:
			return '.'
		case BoxLeft:
			return '['
		case BoxRight:
			return ']'
		default:
			panic(assert.Fail("unexpected item", "item", contents))
		}
	}))
}

func parse(path string, partOne bool) *Warehouse {
	input := strings.Join(utils.MustReadInput(path), "\n")
	parts := strings.Split(input, "\n\n")
	robot := grid.At(-1, -1)
	movements := utils.NewQueue[grid.Coordinate]()
	assert.Equal(2, len(parts), "should have two main parts in the input")

	lines := strings.Split(parts[0], "\n")
	if !partOne {
		for idx, line := range lines {
			line = strings.ReplaceAll(line, "#", "##")
			line = strings.ReplaceAll(line, ".", "..")
			line = strings.ReplaceAll(line, "O", "[]")
			line = strings.ReplaceAll(line, "@", "@.")
			lines[idx] = line
		}
	}

	layout := grid.Parse(lines, func(char rune, x, y int) Item {
		switch char {
		case '#':
			return Wall
		case '@':
			robot = grid.At(x, y)
			return Robot
		case 'O':
			return Box
		case '.':
			return Open
		case '[':
			return BoxLeft
		case ']':
			return BoxRight
		default:
			panic(assert.Fail("unexpected character", "char", char))
		}
	})

	for _, line := range strings.Split(parts[1], "\n") {
		for _, char := range line {
			switch char {
			case '^':
				movements.Enqueue(grid.North)
			case '<':
				movements.Enqueue(grid.West)
			case '>':
				movements.Enqueue(grid.East)
			case 'v':
				movements.Enqueue(grid.South)
			default:
				panic(assert.Fail("unexpected direction", "char", char))
			}
		}
	}

	assert.NotEqual(robot, grid.At(-1, -1), "robot should be set")

	return &Warehouse{
		layout:    layout,
		robot:     robot,
		movements: *movements,
	}
}
