package day16

import (
	"fmt"
	"math"
	"strings"
)

func Part1(input string) (solution int) {
	tiles, start, end := readMaze(input)
	maze := NewGraph(costFunc)
	for tile := range tiles {
		for _, nDir := range []Coord{north, south, east, west} {
			nTile := tile.Add(nDir)
			if tiles[nTile] {
				for _, dir := range []Coord{north, south, east, west} {
					maze.addEdge(Node{tile, dir}, Node{nTile, nDir})
				}
			}
		}
	}

	dist, _ := maze.dijkstra(Node{start, east})

	minDist := math.MaxInt
	for _, dir := range []Coord{north, south, east, west} {
		if d, ok := dist[Node{end, dir}]; ok && d < minDist {
			minDist = d
		}
	}
	return minDist
}

func Part2(input string) (solution int) {
	fmt.Println("No solution yet for day 16, part 2")
	return
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

type Maze Graph[Node]

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
