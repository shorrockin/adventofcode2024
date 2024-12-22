package day12

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/collections"
	"adventofcode2024/pkg/utils/grid"
)

type Group struct {
	value  rune
	coords collections.Set[grid.Coord]
}

func Solve(path string, partOne bool) int {
	plots := parse(path)
	visited := collections.NewSet[grid.Coord]()
	groups := []Group{}

	for coord, node := range plots {
		if visited.Contains(coord) {
			continue
		}

		visited.Add(coord)
		group := Group{node.Contents, collections.NewSetFrom(coord)}
		visit(&plots, coord, &visited, &group)
		groups = append(groups, group)
	}

	return utils.Reduce(groups, 0, func(acc int, group Group) int {
		multiplier := 0
		if partOne {
			multiplier = perimeter(&group)
		} else {
			multiplier = corners(&group)
		}
		return acc + (len(group.coords) * multiplier)
	})
}

func parse(path string) grid.Grid[rune] {
	return grid.Parse(utils.MustReadInput(path), func(value rune, x, y int) rune {
		return value
	})
}

func visit(plots *grid.Grid[rune], source grid.Coord, visited *collections.Set[grid.Coord], group *Group) {
	neighbors := utils.Filter(source.Cardinals(), func(neighbor grid.Coord) bool {
		if visited.Contains(neighbor) {
			return false
		}
		if value, exists := (*plots).GetContents(neighbor); exists {
			return value == group.value
		}
		return false
	})

	for _, neighbor := range neighbors {
		if visited.Contains(neighbor) {
			continue
		}

		visited.Add(neighbor)
		group.coords.Add(neighbor)
		visit(plots, neighbor, visited, group)
	}
}

func perimeter(group *Group) int {
	return utils.Reduce(utils.Keys(group.coords), 0, func(acc int, coord grid.Coord) int {
		perimeter := 4
		neighbors := coord.Cardinals()
		for _, neighbor := range neighbors {
			if group.coords.Contains(neighbor) {
				perimeter--
			}
		}
		return acc + perimeter
	})
}

func corners(group *Group) int {
	corners := 0
	for _, coord := range utils.Keys(group.coords) {
		// checks for corner in the north-east, either an outside corner on an inside
		if !group.coords.Contains(coord.North()) && !group.coords.Contains(coord.East()) {
			corners++
		} else if group.coords.Contains(coord.North()) && group.coords.Contains(coord.East()) && !group.coords.Contains(coord.NorthEast()) {
			corners++
		}

		// checks for corner in the south-east, either an outside corner on an inside
		if !group.coords.Contains(coord.South()) && !group.coords.Contains(coord.East()) {
			corners++
		} else if group.coords.Contains(coord.South()) && group.coords.Contains(coord.East()) && !group.coords.Contains(coord.SouthEast()) {
			corners++
		}

		// checks for corner in the south-west, either an outside corner on an inside
		if !group.coords.Contains(coord.South()) && !group.coords.Contains(coord.West()) {
			corners++
		} else if group.coords.Contains(coord.South()) && group.coords.Contains(coord.West()) && !group.coords.Contains(coord.SouthWest()) {
			corners++
		}

		// checks for corner in the north-west, either an outside corner on an inside
		if !group.coords.Contains(coord.North()) && !group.coords.Contains(coord.West()) {
			corners++
		} else if group.coords.Contains(coord.North()) && group.coords.Contains(coord.West()) && !group.coords.Contains(coord.NorthWest()) {
			corners++
		}
	}
	return corners
}
