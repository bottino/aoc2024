package day20

import (
	"strings"

	"github.com/bottino/aoc2024/vec"
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
			cheatDist := vec.ManhattanDist(x, y)
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

func readInput(input string) (track []vec.Coord, distances map[vec.Coord]int) {
	var start, end vec.Coord
	tiles := make(map[vec.Coord]bool, len(input))
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

	var prev vec.Coord
	curr := start
	track = []vec.Coord{curr}
	distances = map[vec.Coord]int{curr: 0}
	var i int
	for curr != end {
		for _, dir := range vec.AllDirections() {
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
