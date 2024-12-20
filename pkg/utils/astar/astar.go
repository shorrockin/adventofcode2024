package astar

import (
	"adventofcode2024/pkg/utils"
	"container/heap"
)

type Node[T comparable] struct {
	Contents  T
	Priority  float64
	Parent    *Node[T]
	PathDepth int
}

func NewNode[T comparable](contents T, priority float64, parent *Node[T]) *Node[T] {
	depth := 0
	if parent != nil {
		depth = parent.PathDepth + 1
	}
	return &Node[T]{Contents: contents, Priority: priority, Parent: parent, PathDepth: depth}
}

type Preference struct {
	allowBacktrack bool
	includeStart   bool
}

type PriorityQueue[T comparable] []*Node[T]

func (pq PriorityQueue[T]) Len() int           { return len(pq) }
func (pq PriorityQueue[T]) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq PriorityQueue[T]) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*Node[T])
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

type Configurator func(*Preference)

var ExcludeStart Configurator = func(p *Preference) {
	p.includeStart = false
}

var AllowBacktrack Configurator = func(p *Preference) {
	p.allowBacktrack = true
}

func Find[T comparable](start, end T, neighbors func(node *Node[T]) []T, heuristic func(node T, from *Node[T]) float64, configs ...Configurator) []T {
	preferences := &Preference{
		allowBacktrack: false,
		includeStart:   true,
	}
	for _, configurator := range configs {
		configurator(preferences)
	}

	pq := make(PriorityQueue[T], 0)
	visited := utils.NewSet[T]()
	heap.Init(&pq)
	heap.Push(&pq, NewNode(start, 0, nil))

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node[T])

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

			newNode := NewNode(neighbor, heuristic(neighbor, current), current)
			heap.Push(&pq, newNode)
		}
	}

	// no path found, return empty array
	return make([]T, 0)
}
