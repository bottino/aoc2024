package day20

import (
	"strings"
)

func Part1(input string) any {
	return len(findCheats(input, 2, 100))
}

func Part2(input string) any {
	return len(findCheats(input, 20, 100))
}

func findCheats(input string, maxDist int, minCheatSave int) (cheats []int) {
	track, distances := readInput(input)
	for _, x := range track {
		for _, y := range track {
			cheatDist := manDist(x, y)
			if cheatDist <= maxDist {
				cheatSave := distances[y] - distances[x] - cheatDist
				if cheatSave >= minCheatSave {
					cheats = append(cheats, cheatSave)
				}
			}
		}
	}

	return cheats
}

type Coord struct {
	x, y int
}

func (lhs *Coord) Add(rhs Coord) Coord {
	return Coord{lhs.x + rhs.x, lhs.y + rhs.y}
}

func manDist(u Coord, v Coord) int {
	return absInt(u.x-v.x) + absInt(u.y-v.y)
}

func absInt(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

var (
	north = Coord{-1, 0}
	south = Coord{1, 0}
	east  = Coord{0, 1}
	west  = Coord{0, -1}
)

func readInput(input string) (track []Coord, distances map[Coord]int) {
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
	track = []Coord{curr}
	distances = map[Coord]int{curr: 0}
	var i int
	for curr != end {
		for _, dir := range []Coord{north, south, east, west} {
			n := curr.Add(dir)
			if n == prev || !tiles[n] {
				continue
			}

			prev = curr
			curr = n
			i++
			track = append(track, curr)
			distances[curr] = i
		}
	}

	return track, distances
}
