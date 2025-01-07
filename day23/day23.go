package day23

import (
	"slices"
	"strings"

	"github.com/bottino/aoc2024/dsa"
)

func Part1(input string) any {
	g := buildGraph(input)
	cycles := g.FindAllCycles()
	processed := make(map[[3]string]struct{})
	for _, c := range cycles {
		hasTComputer := false
		for _, comp := range c {
			if rune(comp[0]) == 't' {
				hasTComputer = true
			}
		}
		if !hasTComputer {
			continue
		}
		slices.Sort(c)
		processed[[3]string{c[0], c[1], c[2]}] = struct{}{}
	}

	return len(processed)
}

func Part2(input string) any {
	g := buildGraph(input)
	cliques := g.GetCliques()
	var maxClique dsa.Set[string]
	var maxCliqueLen int
	for _, c := range cliques {
		if len(c) > maxCliqueLen {
			maxCliqueLen = len(c)
			maxClique = c
		}
	}

	var password []string
	for k := range maxClique {
		password = append(password, k)
	}
	slices.Sort(password)

	return strings.Join(password, ",")
}

func buildGraph(input string) dsa.Graph[string] {
	g := dsa.NewGraph[string]()
	for _, line := range strings.Split(input, "\n") {
		computers := strings.Split(line, "-")
		g.AddUndirectedEdge(computers[0], computers[1])
	}

	return g
}
