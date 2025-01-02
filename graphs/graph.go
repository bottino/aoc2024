package graphs

import (
	"math"
)

// Graph with integer distance metric
type Graph[T comparable] struct {
	AdjList map[T][]T
	nodes   []T
}

func UnitDist[T comparable](u T, v T) int {
	return 1
}

func (g *Graph[T]) neighbors(node T) []T {
	n, ok := g.AdjList[node]
	if ok {
		return n
	}
	return []T{}
}

func (g *Graph[T]) AddEdge(u T, v T) {
	adjs, ok := g.AdjList[u]
	if ok {
		g.AdjList[u] = append(adjs, v) // ignore duplicate edges
	} else {
		g.AdjList[u] = []T{v}
	}
}

func (g *Graph[T]) Nodes() []T {
	if g.nodes != nil {
		return g.nodes
	}

	nodeMap := make(map[T]bool)
	for k, v := range g.AdjList {
		nodeMap[k] = true
		for _, n := range v {
			nodeMap[n] = true
		}
	}

	var nodes []T
	for k := range nodeMap {
		nodes = append(nodes, k)
	}

	g.nodes = nodes
	return nodes
}

func New[T comparable]() Graph[T] {
	return Graph[T]{make(map[T][]T), nil}
}

func (g *Graph[T]) Dijkstra(source T, distFunc func(T, T) int) (dist map[T]int, prev map[T][]T) {
	dist = make(map[T]int, len(g.Nodes()))
	prev = make(map[T][]T, len(g.Nodes()))

	pq := NewPQueue[T]()
	for _, n := range g.Nodes() {
		dist[n] = math.MaxInt
	}

	dist[source] = 0
	pq.AddWithRank(source, 0)
	for pq.Len() > 0 {
		u := pq.PopMin()
		for _, v := range g.neighbors(u) {
			d := dist[u] + distFunc(u, v)
			// found shortest path yet, resetting prev
			if d <= dist[v] {
				if d < dist[v] {
					dist[v] = d
					prev[v] = []T{}
				}
				prev[v] = append(prev[v], u)
				pq.AddWithRank(v, d)
			}
		}
	}

	return dist, prev
}

func buildPaths[T comparable](curr []T, prev map[T][]T, paths *[][]T) {
	prevs, ok := prev[curr[0]]
	if ok {
		for _, p := range prevs {
			newCurr := append([]T{p}, curr...)
			buildPaths(newCurr, prev, paths)
		}
	} else {
		*paths = append(*paths, curr)
	}
}

func (_ *Graph[T]) GetAllShortestPaths(end T, prev map[T][]T) [][]T {
	var paths [][]T
	curr := []T{end}
	buildPaths(curr, prev, &paths)
	return paths
}
