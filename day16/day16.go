package day16

import (
	"fmt"
	"strings"
)

func Part1(input string) (solution int) {
	tiles, start, end := readMaze(input)
	maze := NewGraph[Tile](func(_ Tile, _ Tile) int { return 1 })
	for tile := range tiles {
		for _, d := range [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			nTile := Tile{tile.x + d[0], tile.y + d[1]}
			if tiles[nTile] {
				maze.addEdge(tile, nTile)
			}
		}
	}

	dist, prev := maze.dijkstra(start)
	dest := prev[end]
	path := []Tile{dest}
	for dest != start {
		dest = prev[dest]
		path = append(path, dest)
	}

	fmt.Println(path)
	return dist[end]
}

func Part2(input string) (solution int) {
	fmt.Println("No solution yet for day 16, part 2")
	return
}

type Maze Graph[Tile]

type Tile struct {
	x, y int
}

func readMaze(input string) (tiles map[Tile]bool, start Tile, end Tile) {
	tiles = make(map[Tile]bool, len(input))
	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			switch char {
			case '#':
				continue
			case '.':
			case 'S':
				start = Tile{i, j}
			case 'E':
				end = Tile{i, j}
			}

			tiles[Tile{i, j}] = true
		}
	}

	return tiles, start, end
}
