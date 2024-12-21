package day21

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/assert"
	"adventofcode2024/pkg/utils/benchmark"
	"adventofcode2024/pkg/utils/grid"
	"fmt"
	"strings"
)

type CacheKey struct {
	from       rune
	to         rune
	iterations int
}

type Keypad struct {
	name   string
	runes  grid.Grid[rune]
	coords map[rune]grid.Coord
	cache  map[CacheKey]string
}

type Candidates struct {
	values []string
	length int
}

var numerics Keypad
var directions Keypad
var bm = benchmark.Start("day21")

func init() {
	create := func(name string, path string) Keypad {
		runes := grid.Parse(utils.MustReadInput(path), grid.ParseRune)
		coords := make(map[rune]grid.Coord)
		for coord, node := range runes {
			coords[node.Contents] = coord
		}
		return Keypad{name, runes, coords, make(map[CacheKey]string)}
	}

	numerics = create("numerics", "numeric.txt")
	directions = create("directions", "direction.txt")
}

func Solve(path string, iterations int) int {
	codes := utils.MustReadInput(path)
	complexity := 0
	iterations--

	for _, code := range codes {
		best := ""
		from := 'A'
		for _, char := range code {
			// fmt.Printf("finding best for %c\n", char)
			best += BestPath(&numerics, from, char, iterations)
			from = char
		}
		bm.Lap(fmt.Sprintf("done %v", code))
		// sequences := ShortestStrings{}
		// sequences.append(KeypadSequencesFor(&numerics, code)...)
		// fmt.Printf("bases: %v\n", sequences.values)
		//
		// for range iterations {
		// 	updated := ShortestStrings{}
		// 	for _, sequence := range sequences.values {
		// 		updated.append(KeypadSequencesFor(&directions, sequence)...)
		// 	}
		// 	sequences = updated
		// }
		// fmt.Printf("best: %s\n", best)
		fmt.Printf("code: %s\n", code)
		fmt.Printf(" len: %d\n", len(best))
		fmt.Println()
		//
		complexity += (len(best) * utils.MustAtoi(code[:3]))
	}

	fmt.Printf("path hits / miss: %v / %v\n", pathHits, pathMiss)
	fmt.Printf("direction hits / miss: %v / %v\n", directionHits, directionMiss)

	return complexity
}

var pathHits int
var pathMiss int
var directionHits int
var directionMiss int

func BestPath(keypad *Keypad, from rune, to rune, iterations int) string {
	/*
		At: A, press: 0 (left once, and press)
		L1: <A
		L2: v<<A >>^A
		L3: <vA<AA>>^A vAA<^A
	*/
	// fmt.Printf("[depth-%v][eval] %c to %c on %v\n", 3-iterations, from, to, keypad.name)
	cacheKey := CacheKey{from, to, iterations}
	if best, ok := keypad.cache[cacheKey]; ok {
		pathHits++
		// fmt.Printf("  cache hit\n")
		return best
	}
	pathMiss++
	if from == to {
		// fmt.Printf("  escape\n")
		return "A"
	}

	candidates := Candidates{}
	candidates.append(KeypadDirections(keypad, from, to)...)

	// for _, c := range candidates.values {
	// 	fmt.Printf("  candidate: %v\n", c)
	// }

	if iterations > 0 {
		updated := Candidates{}

		for _, candidate := range candidates.values {
			from = 'A'
			best := ""
			for _, char := range candidate {
				best += BestPath(&directions, from, char, iterations-1)
				from = char
			}
			updated.append(best)
		}

		candidates = updated
	}

	keypad.cache[cacheKey] = candidates.values[0]
	// fmt.Printf("[depth-%v][result] %c to %c == %v\n", 3-iterations, from, to, candidates.values[0])
	return candidates.values[0]
}

// func KeypadSequencesFor(keypad *Keypad, input string) []string {
// 	sequences := Candidates{}
// 	sequences.append(KeypadDirections(keypad, 'A', rune(input[0]))...)
//
// 	for idx := range sequences.values {
// 		sequences.values[idx] = sequences.values[idx] + "A"
// 	}
// 	from := rune(input[0])
//
// 	for _, char := range input[1:] {
// 		directions := KeypadDirections(keypad, from, char)
// 		updated := Candidates{}
//
// 		for _, sequence := range sequences.values {
// 			for _, direction := range directions {
// 				updated.append(sequence + direction + "A")
// 			}
// 		}
// 		sequences = updated
// 		from = char
// 	}
//
// 	fmt.Printf("returing %v sequences for %s\n", len(sequences.values), input)
// 	return sequences.values
// }

var directionsCache = make(map[CacheKey][]string)

func KeypadDirections(keypad *Keypad, from rune, to rune) []string {
	cacheKey := CacheKey{from, to, 0}
	if directions, ok := directionsCache[cacheKey]; ok {
		directionHits++
		return directions
	}
	directionMiss++

	out := make([]string, 0, 2)
	source := keypad.coords[from]
	dest := keypad.coords[to]

	horizontal := grid.East
	deltaX := utils.AbsInt(source.X - dest.X)
	if source.X > dest.X {
		horizontal = grid.West
	}

	vertical := grid.South
	deltaY := utils.AbsInt(source.Y - dest.Y)
	if source.Y > dest.Y {
		vertical = grid.North
	}

	move := func(d1 grid.Coord, a1 int, d2 grid.Coord, a2 int) {
		at := source
		var builder strings.Builder

		for range a1 {
			at = at.Offset(d1)
			if keypad.runes[at].Contents == '.' {
				return
			}
			builder.WriteRune(DirectionAsRune(d1))
		}
		for range a2 {
			at = at.Offset(d2)
			if keypad.runes[at].Contents == '.' {
				return
			}
			builder.WriteRune(DirectionAsRune(d2))
		}
		builder.WriteRune('A')

		result := builder.String()
		if len(out) == 1 && out[0] == result {
			return
		}
		out = append(out, result)
	}

	move(horizontal, deltaX, vertical, deltaY)
	move(vertical, deltaY, horizontal, deltaX)

	directionsCache[cacheKey] = out
	return out
}

func (c *Candidates) append(values ...string) {
	for _, value := range values {
		if c.length == 0 || len(value) < c.length {
			c.values = []string{value}
			c.length = len(value)
		} else if len(value) == c.length {
			c.values = append(c.values, value)
		}
	}
}

func DirectionAsRune(coord grid.Coord) rune {
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
