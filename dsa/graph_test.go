package dsa

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNodes(t *testing.T) {
	less := func(a string, b string) bool {
		return a < b
	}
	g := NewGraph[string]()
	g.AddEdge("a", "b")
	g.AddEdge("b", "f")

	nodes := g.Nodes()
	if diff := cmp.Diff(NewSet("a", "b", "f"), nodes, cmpopts.SortSlices(less)); diff != "" {
		t.Error(diff)
	}

	g.AddEdge("a", "c")
	g.AddEdge("c", "f")

	nodes = g.Nodes()
	if diff := cmp.Diff(NewSet("a", "b", "c", "f"), nodes, cmpopts.SortSlices(less)); diff != "" {
		t.Error(diff)
	}
}

func TestDijkstra(t *testing.T) {
	g := NewGraph[string]()
	g.AddEdge("a", "b")
	g.AddEdge("a", "c")
	g.AddEdge("b", "f")
	g.AddEdge("c", "f")
	g.AddEdge("a", "d")
	g.AddEdge("d", "e")
	g.AddEdge("e", "f")

	dist, _ := g.Dijkstra("a", UnitDist)

	expDist := map[string]int{"a": 0, "b": 1, "c": 1, "d": 1, "e": 2, "f": 2}
	if diff := cmp.Diff(expDist, dist); diff != "" {
		t.Error(diff)
	}
}

func TestShortestPaths(t *testing.T) {
	g := NewGraph[string]()
	g.AddEdge("a", "b")
	g.AddEdge("a", "c")
	g.AddEdge("b", "f")
	g.AddEdge("c", "f")
	g.AddEdge("a", "d")
	g.AddEdge("d", "e")
	g.AddEdge("e", "f")

	_, prev := g.Dijkstra("a", UnitDist)
	paths := g.GetAllShortestPaths("f", prev)

	expected := [][]string{
		{"a", "b", "f"},
		{"a", "c", "f"},
	}

	if cmp.Diff(expected, paths) != "" {
		t.Errorf("Expected %v, got %v", expected, paths)
	}
}