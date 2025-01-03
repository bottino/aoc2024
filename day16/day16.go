package day16

import (
	"math"
	"strings"

	"github.com/bottino/aoc2024/graphs"
)

func Part1(input string) any {
	maze, start, end := buildMaze(input)
	dist, _ := maze.Dijkstra(Node{start, east}, costFunc)

	minDist := math.MaxInt
	for _, dir := range []Coord{north, south, east, west} {
		if d, ok := dist[Node{end, dir}]; ok && d < minDist {
			minDist = d
		}
	}
	return minDist
}

func Part2(input string) any {
	maze, start, end := buildMaze(input)
	dist, prev := maze.Dijkstra(Node{start, east}, costFunc)

	// A bit dirty; we check for all possible orientation if they have
	// best paths, and only count those
	minDist := math.MaxInt
	for _, dir := range []Coord{north, south, east, west} {
		if d, ok := dist[Node{end, dir}]; ok && d < minDist {
			minDist = d
		}
	}

	var numSeats int
	for _, dir := range []Coord{north, south, east, west} {
		if dist[Node{end, dir}] == minDist {
			numSeats += len(getSeats(Node{end, dir}, prev))
		}
	}

	return numSeats
}

func getSeats(endNode Node, prev map[Node][]Node) map[Coord]bool {
	seats := make(map[Coord]bool)
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

func buildMaze(input string) (maze graphs.Graph[Node], start Coord, end Coord) {
	tiles, start, end := readMaze(input)
	maze = graphs.New[Node]()
	for tile := range tiles {
		for _, nDir := range []Coord{north, south, east, west} {
			nTile := tile.Add(nDir)
			if tiles[nTile] {
				for _, dir := range []Coord{north, south, east, west} {
					maze.AddEdge(Node{tile, dir}, Node{nTile, nDir})
				}
			}
		}
	}

	return maze, start, end
}

type Coord struct {
	x, y int
}

var (
	north = Coord{-1, 0}
	south = Coord{1, 0}
	east  = Coord{0, 1}
	west  = Coord{0, -1}
)

func (lhs *Coord) Add(rhs Coord) Coord {
	return Coord{lhs.x + rhs.x, lhs.y + rhs.y}
}

type Node struct {
	tile Coord
	dir  Coord
}

func costFunc(u Node, v Node) int {
	dot := u.dir.x*v.dir.x + u.dir.y*v.dir.y
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

func readMaze(input string) (tiles map[Coord]bool, start Coord, end Coord) {
	tiles = make(map[Coord]bool, len(input))
	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			switch char {
			case '#':
				continue
			case '.':
			case 'S':
				start = Coord{i, j}
			case 'E':
				end = Coord{i, j}
			}

			tiles[Coord{i, j}] = true
		}
	}

	return tiles, start, end
}
