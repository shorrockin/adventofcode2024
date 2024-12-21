package day21

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/assert"
	"adventofcode2024/pkg/utils/grid"
	"math"
	"strings"
)

var NUMERICS *Keypad
var DIRECTIONS *Keypad

type Keypad struct {
	name   string
	runes  grid.Grid[rune]
	coords map[rune]grid.Coord
	cache  utils.Cache[CacheKey, int]
}

type CacheKey struct {
	from       rune
	to         rune
	iterations int
}

func Solve(path string, iterations int) int {
	codes := utils.MustReadInput(path)
	complexity := 0

	for _, code := range codes {
		complexity += shortestPath(NUMERICS, code, iterations) * utils.MustAtoi(code[:3])
	}

	return complexity
}

func shortestPath(keypad *Keypad, code string, iterations int) int {
	shortest := 0
	source := 'A'

	for _, dest := range code {
		shortest += shortestPathBetween(keypad, source, dest, iterations-1)
		source = dest
	}

	return shortest
}

func shortestPathBetween(keypad *Keypad, from rune, to rune, iterations int) int {
	return keypad.cache.Memoize(CacheKey{from, to, iterations}, func() int {
		if from == to {
			return 1
		}

		candidates := keypadDirections(keypad, from, to)
		shortest := math.MaxInt

		for _, candidate := range candidates {
			current := len(candidate)
			if iterations > 0 {
				current = shortestPath(DIRECTIONS, candidate, iterations)
			}

			if current < shortest {
				shortest = current
			}
		}

		return shortest
	})
}

func keypadDirections(keypad *Keypad, from rune, to rune) []string {
	source := keypad.coords[from]
	dest := keypad.coords[to]

	horizontal := grid.East
	if source.X > dest.X {
		horizontal = grid.West
	}

	vertical := grid.South
	if source.Y > dest.Y {
		vertical = grid.North
	}

	horzAmount := utils.AbsInt(source.X - dest.X)
	vertAmount := utils.AbsInt(source.Y - dest.Y)

	out := make([]string, 0, 2)

	if movement, ok := movementString(keypad, source, horizontal, horzAmount, vertical, vertAmount); ok {
		out = append(out, movement)
	}

	if movement, ok := movementString(keypad, source, vertical, vertAmount, horizontal, horzAmount); ok {
		out = append(out, movement)
	}

	return out
}

func movementString(keypad *Keypad, source grid.Coord, directionOne grid.Coord, amountOne int, directionTwo grid.Coord, amountTwo int) (string, bool) {
	var builder strings.Builder
	at := source

	for range amountOne {
		at = at.Offset(directionOne)
		if keypad.runes[at].Contents == '.' {
			return "", false
		}
		builder.WriteRune(coordToRune(directionOne))
	}

	for range amountTwo {
		at = at.Offset(directionTwo)
		if keypad.runes[at].Contents == '.' {
			return "", false
		}
		builder.WriteRune(coordToRune(directionTwo))
	}

	builder.WriteRune('A')
	return builder.String(), true
}

func coordToRune(coord grid.Coord) rune {
	switch coord {
	case grid.North:
		return '^'
	case grid.East:
		return '>'
	case grid.West:
		return '<'
	case grid.South:
		return 'v'
	default:
		panic(assert.Fail("unknown direction", "direction", coord))
	}
}

func init() {
	create := func(name string, path string) *Keypad {
		runes := grid.Parse(utils.MustReadInput(path), grid.ParseRune)
		coords := make(map[rune]grid.Coord)
		for coord, node := range runes {
			coords[node.Contents] = coord
		}
		return &Keypad{name, runes, coords, make(map[CacheKey]int)}
	}

	NUMERICS = create("numerics", "numeric.txt")
	DIRECTIONS = create("directions", "direction.txt")
}
