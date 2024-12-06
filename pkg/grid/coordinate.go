package grid

import (
	"adventofcode2024/pkg/assert"
	"fmt"
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

func (g Coordinate) String() string {
	if g.label != "" {
		return fmt.Sprintf("%v(x:%d,y:%d)", g.label, g.X, g.Y)
	}
	return fmt.Sprintf("(x:%d,y:%d)", g.X, g.Y)
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
