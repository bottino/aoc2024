package dsa

import (
	"math"
)

// Graph with integer distance metric
type Graph[T comparable] struct {
	adjList map[T][]T
	nodes   []T
	marked  map[T]bool
	cycles  [][]T
}

func UnitDist[T comparable](u T, v T) int {
	return 1
}

func (g *Graph[T]) neighbors(node T) []T {
	n, ok := g.adjList[node]
	if ok {
		return n
	}
	return []T{}
}

func (g *Graph[T]) AddEdge(u T, v T) {
	adjs, ok := g.adjList[u]
	if ok {
		g.adjList[u] = append(adjs, v) // ignore duplicate edges
	} else {
		g.adjList[u] = []T{v}
	}

	g.nodes = nil // reset node caching so they get recalculated
}

func (g *Graph[T]) AddUndirectedEdge(u T, v T) {
	g.AddEdge(u, v)
	g.AddEdge(v, u)
}

func (g *Graph[T]) Nodes() []T {
	if g.nodes != nil {
		return g.nodes
	}

	nodeMap := make(map[T]bool)
	for k, v := range g.adjList {
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

func NewGraph[T comparable]() Graph[T] {
	return Graph[T]{make(map[T][]T), nil, nil, nil}
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

func (g *Graph[T]) FindAllCycles() [][]T {
	g.cycles = nil
	for _, s := range g.Nodes() {
		g.marked = make(map[T]bool, len(g.Nodes()))
		g.recurseCycle(s, s, 0)
	}

	return g.cycles
}

func (g *Graph[T]) recurseCycle(v T, u T, depth int) {
	if depth == 2 {
		return
	}
	g.marked[v] = true
	for _, w := range g.neighbors(v) {
		if !g.marked[w] {
			g.recurseCycle(w, v, depth+1)
		} else if w != u {
			g.cycles = append(g.cycles, []T{u, v, w})
		}
	}
}
