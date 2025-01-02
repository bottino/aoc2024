package graphs

import "math"

// Graph with integer distance metric
type Graph[T comparable] struct {
	adjList map[T][]T
}

func (g *Graph[T]) Neighbors(node T) []T {
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
}

func (g *Graph[T]) NV() int {
	return len(g.adjList)
}

func New[T comparable]() Graph[T] {
	return Graph[T]{make(map[T][]T)}
}

func (g *Graph[T]) Dijkstra(source T, distFunc func(T, T) int) (dist map[T]int, prev map[T][]T) {
	dist = make(map[T]int, g.NV())
	prev = make(map[T][]T, g.NV())

	pq := NewPQueue[T]()
	for v := range g.adjList {
		dist[v] = math.MaxInt
	}

	dist[source] = 0
	pq.AddWithRank(source, 0)
	for pq.Len() > 0 {
		u := pq.PopMin()
		for _, v := range g.Neighbors(u) {
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
