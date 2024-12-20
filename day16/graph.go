package day16

import "math"

// Graph with integer distance metric
type Graph[T comparable] struct {
	adjList  map[T][]T
	distFunc func(T, T) int
}

func (g *Graph[T]) neighbors(node T) []T {
	n, ok := g.adjList[node]
	if ok {
		return n
	}
	return []T{}
}

func (g *Graph[T]) addEdge(u T, v T) {
	adjs, ok := g.adjList[u]
	if ok {
		g.adjList[u] = append(adjs, v) // ignore duplicate edges
	} else {
		g.adjList[u] = []T{v}
	}
}

func (g *Graph[T]) nV() int {
	return len(g.adjList)
}

func NewGraph[T comparable](distFunc func(T, T) int) Graph[T] {
	return Graph[T]{make(map[T][]T), distFunc}
}

func (g *Graph[T]) dijkstra(source T) (dist map[T]int, prev map[T][]T) {
	dist = make(map[T]int, g.nV())
	prev = make(map[T][]T, g.nV())

	pq := NewPQueue[T]()
	for v := range g.adjList {
		dist[v] = math.MaxInt
	}

	dist[source] = 0
	pq.AddWithRank(source, 0)
	for pq.Len() > 0 {
		u := pq.PopMin()
		for _, v := range g.neighbors(u) {
			d := dist[u] + g.distFunc(u, v)
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
