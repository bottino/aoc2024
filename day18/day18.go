package day18

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/bottino/aoc2024/dsa"
	"github.com/bottino/aoc2024/vec"
)

func Part1(input string) any {
	N, M := 71, 71
	maxBytes := 1024
	g := makeGrid(N, M)
	corrBytes := readInput(input)
	for i := 0; i < maxBytes; i++ {
		g.RemoveNode(corrBytes[i])
	}

	dist, _ := g.Dijkstra(vec.Coord{0, 0}, dsa.UnitDist)

	return dist[vec.Coord{N - 1, M - 1}]
}

func Part2(input string) any {
	corrBytes := readInput(input)
	byteBlocking := bisectBlocking(0, len(corrBytes)-1, corrBytes)
	return fmt.Sprintf("%d,%d", byteBlocking.Y, byteBlocking.X)
}

// Find the first blocking byte by dissection
func bisectBlocking(lower int, upper int, corrBytes []vec.Coord) vec.Coord {
	midPoint := lower + (upper-lower)/2

	N, M := 71, 71
	g := makeGrid(N, M)
	for i := 0; i <= midPoint; i++ {
		g.RemoveNode(corrBytes[i])
	}

	dist, _ := g.Dijkstra(vec.Coord{0, 0}, dsa.UnitDist)
	endDist := dist[vec.Coord{N - 1, M - 1}]

	if endDist == math.MaxInt {
		if upper-lower <= 1 {
			return corrBytes[lower]
		}
		return bisectBlocking(lower, midPoint, corrBytes)
	} else {
		if upper-lower <= 1 {
			return corrBytes[upper]
		}
		return bisectBlocking(midPoint, upper, corrBytes)
	}
}

func readInput(input string) (corrBytes []vec.Coord) {
	for _, line := range strings.Split(input, "\n") {
		digits := strings.Split(line, ",")
		x, _ := strconv.Atoi(digits[0])
		y, _ := strconv.Atoi(digits[1])
		// coordinates reverse in our lib convention
		corrBytes = append(corrBytes, vec.Coord{y, x})
	}

	return corrBytes
}

func makeGrid(N, M int) dsa.Graph[vec.Coord] {
	g := dsa.NewGraph[vec.Coord]()
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			c := vec.Coord{i, j}
			for _, d := range vec.AllDirections() {
				n := c.Add(d)
				// check bounds
				if n.X < 0 || n.Y < 0 || n.X >= N || n.Y >= M {
					continue
				}
				g.AddEdge(c, n)
			}
		}
	}

	return g
}
