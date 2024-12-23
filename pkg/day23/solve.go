package day23

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/collections"
	"adventofcode2024/pkg/utils/graph"
	"adventofcode2024/pkg/utils/maps"
	slice "adventofcode2024/pkg/utils/slices"
	"slices"
	"strings"
)

func PartOne(path string) int {
	lan := parse(path)
	processed := collections.NewSet[[3]string]()
	count := 0

	for source, weights := range lan.Edges {
		for _, combo := range slice.Combinations(maps.Keys(weights), 2) {
			if lan.Connected(combo[0], combo[1]) {
				set := [3]string{source, combo[0], combo[1]}
				slices.Sort(set[:])

				if processed.MaybeAdd(set) {
					for _, id := range set {
						if strings.Index(id, "t") == 0 {
							count++
							break
						}
					}
				}
			}
		}
	}

	return count
}

func PartTwo(path string) string {
	lan := parse(path)
	clusters := findClusters(lan, []string{}, lan.Values(), &[][]string{})
	largest := []string{}

	for _, cluster := range *clusters {
		if len(cluster) > len(largest) {
			largest = cluster
		}
	}

	slices.Sort(largest)
	return strings.Join(largest, ",")
}

func findClusters(lan *graph.Graph[string], workset, potentials []string, clusters *[][]string) *[][]string {
	if len(potentials) == 0 {
		cluster := make([]string, len(workset))
		copy(cluster, workset)
		*clusters = append(*clusters, cluster)
		return clusters
	}

	for _, computer := range potentials {
		connections := lan.ConnectedTo(computer)
		findClusters(
			lan,
			append(workset, computer),
			slice.Intersection(potentials, connections),
			clusters,
		)
		potentials = slice.Remove(potentials, computer)
	}

	return clusters
}

func parse(path string) *graph.Graph[string] {
	lan := graph.NewGraph[string]()
	for _, line := range utils.MustReadInput(path) {
		parts := strings.Split(line, "-")
		lan.AddBidirectionalEdge(parts[0], parts[1], 1)
	}
	return lan
}
