package astar

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/priorityqueue"
	"container/heap"
)

type Preference struct {
	allowBacktrack bool
	includeStart   bool
}

type Configurator func(*Preference)

var ExcludeStart Configurator = func(p *Preference) {
	p.includeStart = false
}

var AllowBacktrack Configurator = func(p *Preference) {
	p.allowBacktrack = true
}

func AStar[T comparable](start, end T, neighbors func(node *priorityqueue.Node[T]) []T, heuristic func(node T, from *priorityqueue.Node[T]) float64, configs ...Configurator) []T {
	preferences := &Preference{
		allowBacktrack: false,
		includeStart:   true,
	}
	for _, configurator := range configs {
		configurator(preferences)
	}

	pq := make(priorityqueue.Heap[T], 0)
	visited := utils.NewSet[T]()
	heap.Init(&pq)
	heap.Push(&pq, priorityqueue.NewNode(start, 0, nil))

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*priorityqueue.Node[T])

		if current.Contents == end {
			path := make([]T, 0)
			for current != nil {
				path = append([]T{current.Contents}, path...)
				current = current.Parent
			}

			if !preferences.includeStart && len(path) > 1 {
				return path[1:]
			} else {
				return path
			}
		}

		for _, neighbor := range neighbors(current) {
			if !preferences.allowBacktrack {
				if visited.Contains(neighbor) {
					continue
				}
				visited.Add(neighbor)
			}

			newNode := priorityqueue.NewNode(neighbor, heuristic(neighbor, current), current)
			heap.Push(&pq, newNode)
		}
	}

	// no path found, return empty array
	return make([]T, 0)
}
