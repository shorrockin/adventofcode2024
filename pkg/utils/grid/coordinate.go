package grid

import (
	"adventofcode2024/pkg/assert"
	"fmt"
	"math"
)

type Coordinate struct {
	X     int
	Y     int
	label string
}

func At(x, y int) Coordinate {
	return Coordinate{x, y, ""}
}

func (g Coordinate) WithLabel(label string) Coordinate {
	g.label = label
	return g
}

func (g Coordinate) North() Coordinate     { return At(g.X, g.Y-1) }
func (g Coordinate) East() Coordinate      { return At(g.X+1, g.Y) }
func (g Coordinate) NorthEast() Coordinate { return At(g.X+1, g.Y-1) }
func (g Coordinate) South() Coordinate     { return At(g.X, g.Y+1) }
func (g Coordinate) SouthEast() Coordinate { return At(g.X+1, g.Y+1) }
func (g Coordinate) West() Coordinate      { return At(g.X-1, g.Y) }
func (g Coordinate) SouthWest() Coordinate { return At(g.X-1, g.Y+1) }
func (g Coordinate) NorthWest() Coordinate { return At(g.X-1, g.Y-1) }

func (g Coordinate) Cardinals() []Coordinate {
	return []Coordinate{g.North(), g.East(), g.South(), g.West()}
}
func (g Coordinate) Intercardinals() []Coordinate {
	return []Coordinate{g.NorthEast(), g.SouthEast(), g.SouthWest(), g.NorthWest()}
}
func (g Coordinate) Compass() []Coordinate {
	return []Coordinate{g.North(), g.East(), g.NorthEast(), g.South(), g.SouthEast(), g.West(), g.SouthWest(), g.NorthWest()}
}

func (g Coordinate) Offset(offset Coordinate) Coordinate {
	return At(g.X+offset.X, g.Y+offset.Y)
}

func (g Coordinate) Move(offset Coordinate, times int) Coordinate {
	return At(g.X+offset.X*times, g.Y+offset.Y*times)
}

func (g Coordinate) Bounded(x, y int) Coordinate {
	if g.X >= x {
		g.X = g.X % x
	}
	if g.Y >= y {
		g.Y = g.Y % y
	}
	if g.X < 0 {
		g.X = x + g.X
	}
	if g.Y < 0 {
		g.Y = y + g.Y
	}
	return g
}

func (g Coordinate) TurnRight() Coordinate {
	switch g {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		panic(assert.Fail("couldn't turn right, needs to be a cardinal direction", "dir", g))
	}
}

func (g Coordinate) TurnLeft() Coordinate {
	switch g {
	case North:
		return West
	case East:
		return North
	case South:
		return East
	case West:
		return South
	default:
		panic(assert.Fail("couldn't turn left, needs to be a cardinal direction", "dir", g))
	}
}

func (g Coordinate) Distance(from Coordinate) int {
	return int(math.Abs(float64(g.X-from.X))) + int(math.Abs(float64(g.Y-from.Y)))
}

func (g Coordinate) CoordsInRange(maxRange int) []Coordinate {
	points := make([]Coordinate, 0, CoordsWithinRadius(maxRange))
	for dx := -maxRange; dx <= maxRange; dx++ {
		remaining := maxRange - int(math.Abs(float64(dx)))
		for dy := -remaining; dy <= remaining; dy++ {
			points = append(points, Coordinate{X: g.X + dx, Y: g.Y + dy})
		}
	}

	return points
}

func (g Coordinate) String() string {
	if g.label != "" {
		return fmt.Sprintf("%v(x:%d,y:%d)", g.label, g.X, g.Y)
	}
	return fmt.Sprintf("(x:%d,y:%d)", g.X, g.Y)
}

func CoordsWithinRadius(distance int) int {
	return 2*distance*distance + 2*distance + 1
}

var (
	North      = At(0, -1).WithLabel("North")
	East       = At(1, 0).WithLabel("East")
	South      = At(0, 1).WithLabel("South")
	West       = At(-1, 0).WithLabel("West")
	NorthWest  = At(-1, -1).WithLabel("NorthWest")
	NorthEast  = At(1, -1).WithLabel("NorthEast")
	SouthWest  = At(-1, 1).WithLabel("SouthWest")
	SouthEast  = At(1, 1).WithLabel("SouthEast")
	Directions = []Coordinate{North, East, South, West, NorthWest, NorthEast, SouthWest, SouthEast}
)
