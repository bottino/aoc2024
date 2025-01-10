package dsa

import (
	"math"
)

// Graph with integer distance metric
type Graph[T comparable] struct {
	adjList map[T]Set[T]
	nodes   Set[T]
	marked  map[T]bool
	cycles  [][]T
}

func UnitDist[T comparable](u T, v T) int {
	return 1
}

func (g *Graph[T]) neighbors(node T) Set[T] {
	n, ok := g.adjList[node]
	if ok {
		return n
	}
	return NewSet[T]()
}

func (g *Graph[T]) AddEdge(u T, v T) {
	adjs, ok := g.adjList[u]
	if ok {
		adjs.Add(v)
		g.adjList[u] = adjs // ignore duplicate edges
	} else {
		g.adjList[u] = NewSet(v)
	}

	g.nodes = nil // reset node caching so they get recalculated
}

func (g *Graph[T]) AddUndirectedEdge(u T, v T) {
	g.AddEdge(u, v)
	g.AddEdge(v, u)
}

func (g *Graph[T]) RemoveNode(u T) {
	g.nodes = nil // reset node caching so they get recalculated
	if _, ok := g.adjList[u]; ok {
		delete(g.adjList, u)
	}
	for _, v := range g.adjList {
		v.Remove(u)
	}
}

func (g *Graph[T]) Nodes() Set[T] {
	if g.nodes != nil {
		return g.nodes
	}

	nodeSet := NewSet[T]()
	for k, v := range g.adjList {
		nodeSet.Add(k)
		for n := range v {
			nodeSet.Add(n)
		}
	}

	g.nodes = nodeSet
	return nodeSet
}

func NewGraph[T comparable]() Graph[T] {
	return Graph[T]{make(map[T]Set[T]), nil, nil, nil}
}

func (g *Graph[T]) Dijkstra(source T, distFunc func(T, T) int) (dist map[T]int, prev map[T][]T) {
	dist = make(map[T]int, len(g.Nodes()))
	prev = make(map[T][]T, len(g.Nodes()))

	pq := NewPQueue[T]()
	for n := range g.Nodes() {
		dist[n] = math.MaxInt
	}

	dist[source] = 0
	pq.AddWithRank(source, 0)
	for pq.Len() > 0 {
		u := pq.PopMin()
		for v := range g.neighbors(u) {
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
	for s := range g.Nodes() {
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
	for w := range g.neighbors(v) {
		if !g.marked[w] {
			g.recurseCycle(w, v, depth+1)
		} else if w != u {
			g.cycles = append(g.cycles, []T{u, v, w})
		}
	}
}

// Gets all maximal cliques in the graph using a naive Bron-Kerbosh algo
// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func (g *Graph[T]) GetCliques() []Set[T] {
	var cliques []Set[T]
	g.bronKerbosh(NewSet[T](), g.Nodes(), NewSet[T](), &cliques)
	return cliques
}

func (g *Graph[T]) bronKerbosh(R Set[T], P Set[T], X Set[T], cliques *[]Set[T]) {
	if P.Empty() && X.Empty() {
		if len(R) > 2 {
			*cliques = append(*cliques, R)
		}
		return
	}
	for v := range P {
		g.bronKerbosh(
			R.Union(NewSet(v)),
			P.Intersection(g.neighbors(v)),
			X.Intersection(g.neighbors(v)),
			cliques,
		)
		P.Remove(v)
		X.Add(v)
	}
}
