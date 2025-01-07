package day16

import (
	"math"
	"strings"

	"github.com/bottino/aoc2024/graphs"
	"github.com/bottino/aoc2024/vec"
)

func Part1(input string) any {
	maze, start, end := buildMaze(input)
	dist, _ := maze.Dijkstra(Node{start, vec.East}, costFunc)

	minDist := math.MaxInt
	for _, dir := range vec.AllDirections() {
		if d, ok := dist[Node{end, dir}]; ok && d < minDist {
			minDist = d
		}
	}
	return minDist
}

func Part2(input string) any {
	maze, start, end := buildMaze(input)
	dist, prev := maze.Dijkstra(Node{start, vec.East}, costFunc)

	// A bit dirty; we check for all possible orientation if they have
	// best paths, and only count those
	minDist := math.MaxInt
	for _, dir := range vec.AllDirections() {
		if d, ok := dist[Node{end, dir}]; ok && d < minDist {
			minDist = d
		}
	}

	var numSeats int
	for _, dir := range vec.AllDirections() {
		if dist[Node{end, dir}] == minDist {
			numSeats += len(getSeats(Node{end, dir}, prev))
		}
	}

	return numSeats
}

func getSeats(endNode Node, prev map[Node][]Node) map[vec.Coord]bool {
	seats := make(map[vec.Coord]bool)
	if _, ok := prev[endNode]; !ok {
		return seats
	}

	stack := []Node{endNode}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		seats[current.tile] = true
		for _, p := range prev[current] {
			stack = append(stack, p)
		}
	}

	return seats
}

func buildMaze(input string) (maze graphs.Graph[Node], start vec.Coord, end vec.Coord) {
	tiles, start, end := readMaze(input)
	maze = graphs.New[Node]()
	for tile := range tiles {
		for _, nDir := range vec.AllDirections() {
			nTile := tile.Add(nDir)
			if tiles[nTile] {
				for _, dir := range vec.AllDirections() {
					maze.AddEdge(Node{tile, dir}, Node{nTile, nDir})
				}
			}
		}
	}

	return maze, start, end
}

type Node struct {
	tile vec.Coord
	dir  vec.Coord
}

func costFunc(u Node, v Node) int {
	dot := u.dir.Dot(v.dir)
	switch dot {
	case 0:
		return 1001
	case -1: // can happen on source tile
		return 2001
	case 1:
		return 1
	default:
		panic("Unexpected dot product")
	}
}

func readMaze(input string) (tiles map[vec.Coord]bool, start vec.Coord, end vec.Coord) {
	tiles = make(map[vec.Coord]bool, len(input))
	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			switch char {
			case '#':
				continue
			case '.':
			case 'S':
				start = vec.Coord{i, j}
			case 'E':
				end = vec.Coord{i, j}
			}

			tiles[vec.Coord{i, j}] = true
		}
	}

	return tiles, start, end
}
