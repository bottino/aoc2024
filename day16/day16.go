package day16

import (
	"fmt"
	"strings"
)

func Part1(input string) (solution int) {
	tiles, start, end := readMaze(input)
	maze := NewGraph(costFunc)
	for tile := range tiles {
		for _, d := range []Coord{north, south, east, west} {
			nTile := tile.Add(d)
			if tiles[nTile] {
				maze.addEdge(tile, nTile)
			}
		}
	}

	dist, prev := maze.dijkstra(start)

	dest := prev[end]
	path := []Coord{dest}
	for dest != start {
		dest = prev[dest]
		path = append(path, dest)
	}

	// fmt.Println(path)
	return dist[end]
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

type Maze Graph[Coord]

func costFunc(b Coord, c Coord, a Coord) int {
	ab := Coord{b.x - a.x, b.y - a.y}
	bc := Coord{c.x - b.x, c.y - b.y}

	// Handle no previous
	var defaultCoord Coord
	if a == defaultCoord {
		ab = east
	}

	dot := ab.x*bc.x + ab.y*bc.y

	if a == defaultCoord {
		fmt.Println(a, b, c, dot)
	}
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
