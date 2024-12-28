package day20

import (
	"fmt"
	"strings"
)

func Part1(input string) any {
	return 0
}

func Part2(input string) any {
	fmt.Println("No solution yet for day 20, part 2")
	return 0
}

type Coord struct {
	x, y int
}

func (lhs *Coord) Add(rhs Coord) Coord {
	return Coord{lhs.x + rhs.x, lhs.y + rhs.y}
}

var (
	north = Coord{-1, 0}
	south = Coord{1, 0}
	east  = Coord{0, 1}
	west  = Coord{0, -1}
)

func readInput(input string) []Coord {
	var start, end Coord
	tiles := make(map[Coord]bool, len(input))
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

	var prev Coord
	curr := start
	track := []Coord{curr}
	for curr != end {
		for _, dir := range []Coord{north, south, east, west} {
			n := curr.Add(dir)
			if n == prev || !tiles[n] {
				continue
			}

			prev = curr
			curr = n
			track = append(track, curr)
		}
	}

	return track
}
